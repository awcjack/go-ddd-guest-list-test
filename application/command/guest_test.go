package command_test

import (
	"context"
	"testing"

	"github.com/awcjack/getground-backend-assignment/application/command"
	"github.com/awcjack/getground-backend-assignment/domain/guest"
	"github.com/awcjack/getground-backend-assignment/domain/table"
	"github.com/sirupsen/logrus"
)

func TestAddGuest(t *testing.T) {
	dep := newGuestDependencies()
	table, _ := table.NewTable(0, 1, 1, 1)
	dep.addGuestHandler.Handle(context.Background(), guest.NewGuest("John", 1, &table))

	if dep.Repository.CalledAddGuestTime != 1 ||
		dep.Repository.CalledListArrivedGuestsTime != 0 ||
		dep.Repository.CalledListGuestsTime != 0 ||
		dep.Repository.CalledUpdateGuestTime != 0 {
		t.Errorf("Expected only call addGuest once, but got %v", dep.Repository)
	}
}

func TestCheckInGuest(t *testing.T) {
	dep := newGuestDependencies()
	dep.checkInGuestHandler.Handle(context.Background(), "John", 1)

	if dep.Repository.CalledAddGuestTime != 0 ||
		dep.Repository.CalledListArrivedGuestsTime != 0 ||
		dep.Repository.CalledListGuestsTime != 0 ||
		dep.Repository.CalledUpdateGuestTime != 1 {
		t.Errorf("Expected only call updateGuest once, but got %v", dep.Repository)
	}
}

func TestCheckoutGuest(t *testing.T) {
	dep := newGuestDependencies()
	dep.checkOutGuestHandler.Handle(context.Background(), "John")

	if dep.Repository.CalledAddGuestTime != 0 ||
		dep.Repository.CalledListArrivedGuestsTime != 0 ||
		dep.Repository.CalledListGuestsTime != 0 ||
		dep.Repository.CalledUpdateGuestTime != 1 {
		t.Errorf("Expected only call updateGuest once, but got %v", dep.Repository)
	}
}

// struct that keep the mocked repository implementation and handler
type guestDependencies struct {
	Repository           *guestRepoMock
	addGuestHandler      *command.AddGuestHandler
	checkInGuestHandler  *command.CheckInGuestHandler
	checkOutGuestHandler *command.CheckOutGuestHandler
}

// guestDependencies constructor
func newGuestDependencies() guestDependencies {
	repository := &guestRepoMock{}
	logger := logrus.NewEntry(logrus.StandardLogger())

	return guestDependencies{
		Repository:           repository,
		addGuestHandler:      command.NewAddGuestHandler(repository, logger),
		checkInGuestHandler:  command.NewCheckInGuestHandler(repository, logger),
		checkOutGuestHandler: command.NewCheckOutGuestHandler(repository, logger),
	}
}

// Mock repository implementation for tracking the handler called correct function
type guestRepoMock struct {
	CalledAddGuestTime          int
	CalledListGuestsTime        int
	CalledListArrivedGuestsTime int
	CalledUpdateGuestTime       int
}

func (g *guestRepoMock) AddGuest(ctx context.Context, guest guest.Guest) error {
	g.CalledAddGuestTime++
	return nil
}

func (g *guestRepoMock) ListGuests(ctx context.Context) ([]guest.Guest, error) {
	g.CalledListGuestsTime++
	return nil, nil
}

func (g *guestRepoMock) ListArrivedGuests(ctx context.Context) ([]guest.Guest, error) {
	g.CalledListArrivedGuestsTime++
	return nil, nil
}

func (g *guestRepoMock) UpdateGuest(ctx context.Context, name string, updateFn func(g *guest.Guest) (*guest.Guest, error)) error {
	g.CalledUpdateGuestTime++
	return nil
}
