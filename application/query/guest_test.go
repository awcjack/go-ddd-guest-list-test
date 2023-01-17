package query_test

import (
	"context"
	"testing"

	"github.com/awcjack/getground-backend-assignment/application/query"
	"github.com/awcjack/getground-backend-assignment/domain/guest"
	"github.com/sirupsen/logrus"
)

func TestListGuest(t *testing.T) {
	dep := newGuestDependencies()
	dep.listGuestHandler.Handle(context.Background())

	if dep.Repository.CalledAddGuestTime != 0 ||
		dep.Repository.CalledListArrivedGuestsTime != 0 ||
		dep.Repository.CalledListGuestsTime != 1 ||
		dep.Repository.CalledUpdateGuestTime != 0 {
		t.Errorf("Expected only call list guest once, but got %v", dep.Repository)
	}
}

func TestListArrivedGuest(t *testing.T) {
	dep := newGuestDependencies()
	dep.listArrivedGuestHandler.Handle(context.Background())

	if dep.Repository.CalledAddGuestTime != 0 ||
		dep.Repository.CalledListArrivedGuestsTime != 1 ||
		dep.Repository.CalledListGuestsTime != 0 ||
		dep.Repository.CalledUpdateGuestTime != 0 {
		t.Errorf("Expected only call list arrived guest once, but got %v", dep.Repository)
	}
}

// struct that keep the mocked repository implementation and handler
type guestDependencies struct {
	Repository              *guestRepoMock
	listGuestHandler        *query.ListGuestHandler
	listArrivedGuestHandler *query.ListArrivedGuestHandler
}

// guestDependencies constructor
func newGuestDependencies() guestDependencies {
	repository := &guestRepoMock{}
	logger := logrus.NewEntry(logrus.StandardLogger())

	return guestDependencies{
		Repository:              repository,
		listGuestHandler:        query.NewListGuestHandler(repository, logger),
		listArrivedGuestHandler: query.NewListArrivedGuestHandler(repository, logger),
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
