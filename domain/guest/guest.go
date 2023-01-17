package guest

import (
	"errors"
	"time"

	"github.com/awcjack/getground-backend-assignment/domain/table"
)

// All fields are private to prevent updated in somewhere (Updating the value is allowed iff setter is provided below)
type Guest struct {
	// Guest name (unique identifier)
	name string
	// Expected guest number when adding to the guest list
	guestNumber int
	// Corresponding table for guest
	table *table.Table
	// Guest arrival status
	arrived bool
	// Guest number when arrived
	arrivedNumber int
	// Arrived time
	arrivedTime time.Time
}

var (
	ErrGuestAlreadyArrived = errors.New("guest already arrived")
	ErrGuestNotArrived     = errors.New("guest doesn't arrived yet")
)

// Constructor to create guest entity
func NewGuest(name string, guestNumber int, table *table.Table) Guest {
	return Guest{
		name:        name,
		guestNumber: guestNumber,
		table:       table,
	}
}

// Retrieve Guest name
func (g Guest) Name() string {
	return g.name
}

// Retrieve Guest Table info
func (g Guest) Table() *table.Table {
	return g.table
}

// Retrieve Expected guest number
func (g Guest) GuestNumber() int {
	return g.guestNumber
}

// Retrieve Guest arrival status
func (g Guest) Arrived() bool {
	return g.arrived
}

// Retrieve Guest number when arrived
func (g Guest) ArrivedNumber() int {
	return g.arrivedNumber
}

// Retrieve arrived time
func (g Guest) ArrivedTime() time.Time {
	return g.arrivedTime
}

// Checkin guest
func (g *Guest) Arrive(guestNumber int) error {
	if g.arrived {
		return ErrGuestAlreadyArrived
	}

	err := g.table.Checkin(guestNumber)
	if err != nil {
		return err
	}

	g.arrived = true
	g.arrivedNumber = guestNumber
	g.arrivedTime = time.Now()

	return nil
}

// arrival status setter (for testing)
func (g *Guest) SetArrived(arrived bool) {
	g.arrived = arrived
}

// arrival guest number setter (for testing)
func (g *Guest) SetArriveNumber(guestNumber int) {
	g.arrivedNumber = guestNumber
}

// arrival time setter (for testing)
func (g *Guest) SetArrivedTime(time time.Time) {
	g.arrivedTime = time
}

// Checkout guest
func (g *Guest) Leave() error {
	if !g.arrived {
		return ErrGuestNotArrived
	}

	g.table.Checkout(g.arrivedNumber)

	g.arrived = false
	g.arrivedNumber = 0
	g.arrivedTime = time.Time{}

	return nil
}

// Empty Guest check
func (g Guest) IsZero() bool {
	return g.name == "" && g.guestNumber == 0 && g.table == nil && !g.arrived && g.arrivedNumber == 0 && g.arrivedTime.IsZero()
}
