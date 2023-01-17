package datastore

import (
	"context"
	"time"

	"github.com/awcjack/getground-backend-assignment/config"
	"github.com/awcjack/getground-backend-assignment/domain/guest"
	"github.com/awcjack/getground-backend-assignment/domain/table"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// repository implementation in MySQL storage
type MySQLRepository struct {
	db     *sqlx.DB
	logger logger
}

// MySQLRepository constructor
func NewMySQLRepository(db *sqlx.DB, logger logger) *MySQLRepository {
	if db == nil {
		logger.Panicf("missing db")
	}

	return &MySQLRepository{
		db:     db,
		logger: logger,
	}
}

// table entity when interact with MySQL
type mysqlTable struct {
	Id             int `db:"id"`
	Capacity       int `db:"capacity"`
	AvailableSpace int `db:"availableSpace"`
	AvailableSeat  int `db:"availableSeat"`
}

// guest entity when interact with MySQL
type mysqlGuest struct {
	Name          string `db:"name"`
	GuestNumber   int    `db:"guest_number"`
	Table         int    `db:"table_id"`
	Arrived       bool   `db:"arrived"`
	ArrivedNumber int    `db:"arrived_number"`
	ArrivedTime   string `db:"time_arrived"`
}

// Add table to persistent layer implementation
func (m *MySQLRepository) AddTable(ctx context.Context, t *table.Table) error {
	tx, err := m.db.Beginx()
	if err != nil {
		return errors.Wrap(err, "unable to start transaction")
	}

	defer func() {
		err = m.finishTransaction(err, tx)
	}()

	dbTable := mysqlTable{
		Capacity: t.Capacity(),
	}

	// Insert table to tables table in MySQL
	result, err := tx.NamedExec(`
		INSERT INTO
			tables (capacity)
		VALUES
			(:capacity)
	`, dbTable)
	if err != nil {
		return err
	}

	// set id after insertingt table to MySQL
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	// set Id once table persist in db
	t.SetId(int(id))
	return nil
}

// Get table to persistent layer implementation
func (m *MySQLRepository) GetTable(ctx context.Context, id int) (*table.Table, error) {
	tx, err := m.db.Beginx()
	if err != nil {
		return nil, errors.Wrap(err, "unable to start transaction")
	}

	defer func() {
		err = m.finishTransaction(err, tx)
	}()

	dbTable := mysqlTable{}

	// get table from MySQL and calculate the available space and available seat in MySQL
	// set space and seat to capacity if no guest using that table yet to prevent null case
	err = tx.GetContext(ctx, &dbTable, "SELECT id, capacity, IFNULL(((capacity) - SUM(guest_number)), capacity) as availableSpace, IFNULL(((capacity) - SUM(arrived_number)), capacity) as availableSeat FROM `tables` LEFT JOIN `guests` on  `tables`.id = `guests`.table_id WHERE `id` = ? GROUP BY id", id)
	if err != nil {
		return nil, errors.New("Table not found from db")
	}

	// convert MySQL table to domain table entity
	domainTable, err := table.NewTable(dbTable.Id, dbTable.Capacity, dbTable.AvailableSpace, dbTable.AvailableSeat)
	if err != nil {
		return nil, err
	}

	return &domainTable, nil
}

// List table to persistent layer implementation
func (m *MySQLRepository) ListTable(ctx context.Context) ([]table.Table, error) {
	tx, err := m.db.Beginx()
	if err != nil {
		return nil, errors.Wrap(err, "unable to start transaction")
	}

	defer func() {
		err = m.finishTransaction(err, tx)
	}()

	domainTables := make([]table.Table, 0)

	// get all tables from MySQL and calculate the available space and available seat in MySQL
	rows, err := tx.QueryxContext(ctx, "SELECT id, capacity, IFNULL(((capacity) - SUM(guest_number)), capacity) as availableSpace, IFNULL(((capacity) - SUM(arrived_number)), capacity) as availableSeat FROM `tables` LEFT JOIN `guests` on  `tables`.id = `guests`.table_id GROUP BY id")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		dbTable := mysqlTable{}
		err := rows.Scan(&dbTable.Id, &dbTable.Capacity, &dbTable.AvailableSpace, &dbTable.AvailableSeat)
		if err != nil {
			return nil, err
		}
		domainTable, err := table.NewTable(dbTable.Id, dbTable.Capacity, dbTable.AvailableSpace, dbTable.AvailableSeat)
		if err != nil {
			return nil, err
		}
		domainTables = append(domainTables, domainTable)
	}

	return domainTables, nil
}

// Add guest to persistent layer implementation
func (m *MySQLRepository) AddGuest(ctx context.Context, g guest.Guest) error {
	tx, err := m.db.Beginx()
	if err != nil {
		return errors.Wrap(err, "unable to start transaction")
	}

	defer func() {
		err = m.finishTransaction(err, tx)
	}()

	err = g.Table().AllocateSeat(g.GuestNumber())
	if err != nil {
		return err
	}

	dbGuest := mysqlGuest{
		Name:          g.Name(),
		GuestNumber:   g.GuestNumber(),
		Table:         g.Table().Id(),
		Arrived:       g.Arrived(),
		ArrivedNumber: g.ArrivedNumber(),
		ArrivedTime:   g.ArrivedTime().Format("2006-01-02 15:04:05"),
	}

	// Inserting guest to MySQL database
	_, err = tx.NamedExec(`
		INSERT INTO
			guests (name, guest_number, table_id, arrived, arrived_number, time_arrived)
		VALUES
			(:name, :guest_number, :table_id, :arrived, :arrived_number, :time_arrived)
	`, dbGuest)
	if err != nil {
		return ErrGuestAlreadyExist
	}

	return nil
}

// List guest to persistent layer implementation
func (m *MySQLRepository) ListGuests(ctx context.Context) ([]guest.Guest, error) {
	tx, err := m.db.Beginx()
	if err != nil {
		return nil, errors.Wrap(err, "unable to start transaction")
	}

	defer func() {
		err = m.finishTransaction(err, tx)
	}()

	domainGuests := make([]guest.Guest, 0)

	// Get all guest from MySQL
	rows, err := tx.QueryxContext(ctx, "SELECT name, guest_number, table_id FROM `guests`")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		dbGuest := mysqlGuest{}
		err := rows.Scan(&dbGuest.Name, &dbGuest.GuestNumber, &dbGuest.Table)
		if err != nil {
			return nil, err
		}
		fakeTable, _ := table.NewTable(dbGuest.Table, 0, 0, 0)
		domainGuest := guest.NewGuest(dbGuest.Name, dbGuest.GuestNumber-1, &fakeTable)
		domainGuests = append(domainGuests, domainGuest)
	}

	return domainGuests, nil
}

// List arrived guest to persistent layer implementation
func (m *MySQLRepository) ListArrivedGuests(ctx context.Context) ([]guest.Guest, error) {
	tx, err := m.db.Beginx()
	if err != nil {
		return nil, errors.Wrap(err, "unable to start transaction")
	}

	defer func() {
		err = m.finishTransaction(err, tx)
	}()

	domainGuests := make([]guest.Guest, 0)

	// Get all arrived guest from MySQL
	rows, err := tx.QueryxContext(ctx, "SELECT name, arrived_number, time_arrived FROM `guests` where arrived = true")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		dbGuest := mysqlGuest{}
		err := rows.Scan(&dbGuest.Name, &dbGuest.ArrivedNumber, &dbGuest.ArrivedTime)
		if err != nil {
			return nil, err
		}
		fakeTable, _ := table.NewTable(dbGuest.Table, 0, 0, 0)
		domainGuest := guest.NewGuest(dbGuest.Name, dbGuest.GuestNumber-1, &fakeTable)
		domainGuest.SetArriveNumber(dbGuest.ArrivedNumber)
		arrivedTime, err := time.Parse("2006-01-02 15:04:05", dbGuest.ArrivedTime)
		if err != nil {
			return nil, err
		}
		domainGuest.SetArrivedTime(arrivedTime)
		domainGuests = append(domainGuests, domainGuest)
	}

	return domainGuests, nil
}

// update guest to persistent layer implementation
func (m *MySQLRepository) UpdateGuest(ctx context.Context, name string, updateFn func(guest *guest.Guest) (*guest.Guest, error)) error {
	tx, err := m.db.Beginx()
	if err != nil {
		return errors.Wrap(err, "unable to start transaction")
	}

	defer func() {
		err = m.finishTransaction(err, tx)
	}()

	dbGuest := mysqlGuest{}

	// Getting guest from MySQL
	err = tx.GetContext(ctx, &dbGuest, "SELECT name, guest_number, table_id, arrived, arrived_number, time_arrived FROM `guests` where `name` = ?", name)
	if err != nil {
		return ErrGuestNotExist
	}

	dbTable := mysqlTable{}

	// Getting corresponding table for guest from MySQL
	err = tx.GetContext(ctx, &dbTable, "SELECT id, capacity, IFNULL(((capacity) - SUM(guest_number)), capacity) as availableSpace, IFNULL(((capacity) - SUM(arrived_number)), capacity) as availableSeat FROM `tables` LEFT JOIN `guests` on  `tables`.id = `guests`.table_id WHERE `id` = ? GROUP BY id", dbGuest.Table)
	if err != nil {
		return ErrTableNotExist
	}

	domainTable, err := table.NewTable(dbTable.Id, dbTable.Capacity, dbTable.AvailableSpace, dbTable.AvailableSeat)
	if err != nil {
		return err
	}

	currentGuest := guest.NewGuest(name, dbGuest.GuestNumber, &domainTable)

	currentGuest.SetArrived(dbGuest.Arrived)
	currentGuest.SetArriveNumber(dbGuest.ArrivedNumber)
	arrivedTime, err := time.Parse("2006-01-02 15:04:05", dbGuest.ArrivedTime)
	if err != nil {
		return err
	}
	currentGuest.SetArrivedTime(arrivedTime)

	_, err = updateFn(&currentGuest)
	if err != nil {
		return err
	}

	// Update the updated guest in MySQL
	_, err = tx.NamedExec(`
		UPDATE guests
		SET arrived=:arrived, arrived_number=:arrived_number, time_arrived=:time_arrived
		WHERE name=:name
	`, mysqlGuest{
		Name:          currentGuest.Name(),
		Arrived:       currentGuest.Arrived(),
		ArrivedNumber: currentGuest.ArrivedNumber(),
		ArrivedTime:   currentGuest.ArrivedTime().Format("2006-01-02 15:04:05"), // format to MySQL datetime format
	})
	if err != nil {
		return err
	}

	return nil
}

// MySQL finish transaction operation (rollback if failure and commit if no error)
func (m MySQLRepository) finishTransaction(err error, tx *sqlx.Tx) error {
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return errors.Wrap(err, rbErr.Error())
		}

		return err
	}
	if commitErr := tx.Commit(); commitErr != nil {
		return errors.Wrap(err, "failed to commit transaction")
	}
	return err
}

// Start MySQL connection
func NewMYSQLConnection(c config.DatabaseConfig) (*sqlx.DB, error) {
	config := mysql.NewConfig()

	config.Net = "tcp"
	config.Addr = c.Addr
	config.User = c.User
	config.Passwd = c.Password
	config.DBName = c.DBName

	db, err := sqlx.Connect("mysql", config.FormatDSN())
	if err != nil {
		return nil, err
	}

	return db, nil
}
