package datastore

import (
	"context"

	"github.com/awcjack/getground-backend-assignment/domain/guest"
	"github.com/awcjack/getground-backend-assignment/domain/table"
)

// repository implementation in memory storage
type MemoryRepository struct {
	tables []*table.Table
	guests []guest.Guest
	logger logger
}

// MemoryRepository constructor
func NewMemoryRepository(logger logger) *MemoryRepository {
	return &MemoryRepository{
		tables: make([]*table.Table, 0),
		guests: make([]guest.Guest, 0),
		logger: logger,
	}
}

// Add table to persistent layer implementation
func (m *MemoryRepository) AddTable(ctx context.Context, t *table.Table) error {
	t.SetId(len(m.tables))
	m.tables = append(m.tables, t)
	m.logger.Debugf("Table after adding table", m.tables)
	return nil
}

// Get table to persistent layer implementation
func (m *MemoryRepository) GetTable(ctx context.Context, id int) (*table.Table, error) {
	m.logger.Debugf("Tables list when get table ", id)
	if id >= len(m.tables) {
		return nil, ErrTableNotExist
	}
	return m.tables[id], nil
}

// List table to persistent layer implementation
func (m *MemoryRepository) ListTable(ctx context.Context) ([]table.Table, error) {
	m.logger.Debugf("Tables list when listing table ", m.tables)
	tables := make([]table.Table, len(m.tables))
	for i, t := range m.tables {
		// inserted record should never throw error
		tables[i], _ = table.NewTable(i, t.Capacity(), t.AvailableSpace(), t.AvailableSeat())
	}
	return tables, nil
}

// Add guest to persistent layer implementation
func (m *MemoryRepository) AddGuest(ctx context.Context, g guest.Guest) error {
	for _, storedG := range m.guests {
		if g.Name() == storedG.Name() {
			return ErrGuestAlreadyExist
		}
	}

	err := g.Table().AllocateSeat(g.GuestNumber())
	if err != nil {
		return err
	}

	m.guests = append(m.guests, g)
	m.logger.Debugf("Guest list after adding guest ", m.guests)
	return nil
}

// List guest to persistent layer implementation
func (m *MemoryRepository) ListGuests(ctx context.Context) ([]guest.Guest, error) {
	m.logger.Debugf("Guest list when listing guest ", m.guests)
	return m.guests, nil
}

// List arrived guest to persistent layer implementation
func (m *MemoryRepository) ListArrivedGuests(ctx context.Context) ([]guest.Guest, error) {
	m.logger.Debugf("Guest list when listing arrived guest ", m.guests)
	var arrivedGuests []guest.Guest
	for _, g := range m.guests {
		if g.Arrived() {
			arrivedGuests = append(arrivedGuests, g)
		}
	}
	return arrivedGuests, nil
}

// update guest to persistent layer implementation
func (m *MemoryRepository) UpdateGuest(ctx context.Context, name string, updateFn func(guest *guest.Guest) (*guest.Guest, error)) error {
	m.logger.Debugf("Guest list before updating guest ", m.guests)
	var currentGuest guest.Guest
	var guestIndex int
	for i, g := range m.guests {
		if g.Name() == name {
			currentGuest = g
			guestIndex = i
		}
	}
	if currentGuest.IsZero() {
		return ErrGuestNotExist
	}

	_, err := updateFn(&currentGuest)
	m.guests[guestIndex] = currentGuest
	if err != nil {
		return err
	}

	m.logger.Debugf("Guest list after updating guest ", m.guests)
	return nil
}
