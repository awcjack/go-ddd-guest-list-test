package table

import (
	"errors"
)

type Table struct {
	// Table Id (UUID/GUID is more preferred in DDD) but here is using the RDBMS auto increment id (Only set after getting the id)
	id int
	// Table capacity
	capacity int
	// Available Space in table (expected space from expected guest number)
	availableSpace int
	// Available Seat after guest arrive
	availableSeat int
}

var (
	ErrNegativeCapacity = errors.New("negative capacity")
	ErrNotEnoughSeat    = errors.New("not enough seat")
)

// Table constructor
func NewTable(id, capacity, availableSpace, availableSeat int) (Table, error) {
	// Table with negative capacity is not allowed
	if capacity < 0 {
		return Table{}, ErrNegativeCapacity
	}

	return Table{
		id:             id,
		capacity:       capacity,
		availableSpace: availableSpace,
		availableSeat:  availableSeat,
	}, nil
}

// Retrieve Table id
func (t Table) Id() int {
	return t.id
}

// Table Id setter, Update id once the entity persist in db
func (t *Table) SetId(id int) {
	t.id = id
}

// Retrieve Table capacity
func (t Table) Capacity() int {
	return t.capacity
}

// Retrieve Table available space
func (t Table) AvailableSpace() int {
	return t.availableSpace
}

// Allocate guest to table iff table have enough space
func (t *Table) AllocateSeat(guestNumber int) error {
	if t.availableSpace >= guestNumber {
		t.availableSpace -= guestNumber
		return nil
	}

	return ErrNotEnoughSeat
}

// Retrieve Table available seat
func (t Table) AvailableSeat() int {
	return t.availableSeat
}

// Allocate guest to table when guest arrive iff table have enough seat
func (t *Table) Checkin(guestNumber int) error {
	if t.availableSeat >= guestNumber {
		t.availableSeat -= guestNumber
		return nil
	}

	return ErrNotEnoughSeat
}

// Release seat for guest when guest leave
func (t *Table) Checkout(guestNumber int) {
	t.availableSeat += guestNumber
}
