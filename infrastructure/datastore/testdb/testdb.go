//go:build integration
// +build integration

// Integration test sample data

package testdb

import (
	"time"

	"github.com/jmoiron/sqlx"
)

// Integration test default table data
func SeedTables(db *sqlx.DB) error {
	tables := []struct {
		id       int
		capacity int
	}{
		{
			id:       1,
			capacity: 1,
		},
		{
			id:       2,
			capacity: 10,
		},
		{
			id:       3,
			capacity: 5,
		},
	}

	for _, table := range tables {
		stmt, err := db.Prepare("INSERT INTO tables(id, capacity) VALUES (?, ?)")
		if err != nil {
			return err
		}

		_, err = stmt.Exec(table.id, table.capacity)
		if err != nil {
			return err
		}
	}
	return nil
}

// Integration test default guest data
func SeedGuests(db *sqlx.DB) error {
	guests := []struct {
		name          string
		guestNumber   int
		tableId       int
		arrived       bool
		arrivedNumber int
		arrivedTime   time.Time
	}{
		{
			name:          "John",
			guestNumber:   1,
			tableId:       1,
			arrived:       false,
			arrivedNumber: 0,
			arrivedTime:   time.Time{},
		},
		{
			name:          "John1",
			guestNumber:   3,
			tableId:       2,
			arrived:       true,
			arrivedNumber: 4,
			arrivedTime:   time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
		},
		{
			name:          "John2",
			guestNumber:   5,
			tableId:       3,
			arrived:       true,
			arrivedNumber: 5,
			arrivedTime:   time.Date(2009, time.November, 11, 23, 0, 0, 0, time.UTC),
		},
		{
			name:          "John3",
			guestNumber:   1,
			tableId:       2,
			arrived:       false,
			arrivedNumber: 0,
			arrivedTime:   time.Time{},
		},
	}

	for _, guest := range guests {
		stmt, err := db.Prepare("INSERT INTO guests(name, guest_number, table_id, arrived, arrived_number, time_arrived) VALUES (?, ?, ?, ?, ?, ?)")
		if err != nil {
			return err
		}

		_, err = stmt.Exec(guest.name, guest.guestNumber, guest.tableId, guest.arrived, guest.arrivedNumber, guest.arrivedTime.Format("2006-01-02 15:04:05"))
		if err != nil {
			return err
		}
	}
	return nil
}

// Integration test clear table data
func CleanTables(db *sqlx.DB) error {
	_, err := db.Exec("ALTER TABLE `guests` DROP FOREIGN KEY table_table_id")
	if err != nil {
		return err
	}
	_, err = db.Exec("TRUNCATE table tables")
	if err != nil {
		return err
	}
	_, err = db.Exec("ALTER TABLE `guests` ADD CONSTRAINT table_table_id FOREIGN KEY (`table_id`) REFERENCES `tables`(`id`)")
	if err != nil {
		return err
	}
	return nil
}

// Integration test clear table data
func CleanGuests(db *sqlx.DB) error {
	_, err := db.Exec("TRUNCATE table guests")
	if err != nil {
		return err
	}
	return nil
}
