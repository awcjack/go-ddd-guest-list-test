package query

import (
	"context"

	"github.com/awcjack/getground-backend-assignment/domain/guest"
)

// Handler for listing guest
type ListGuestHandler struct {
	guestRepo guest.Repository
	logger    logger
}

// ListGuestHandler constructor
func NewListGuestHandler(guestRepo guest.Repository, logger logger) *ListGuestHandler {
	return &ListGuestHandler{
		guestRepo: guestRepo,
		logger:    logger,
	}
}

// Handle list guest operation
func (l ListGuestHandler) Handle(ctx context.Context) ([]guest.Guest, error) {
	guests, err := l.guestRepo.ListGuests(ctx)
	l.logger.Debugf("List guest %v", guests)
	return guests, err
}

// Handler for listing arrived guest
type ListArrivedGuestHandler struct {
	guestRepo guest.Repository
	logger    logger
}

// ListArrivedGuestHandler constructor
func NewListArrivedGuestHandler(guestRepo guest.Repository, logger logger) *ListArrivedGuestHandler {
	return &ListArrivedGuestHandler{
		guestRepo: guestRepo,
		logger:    logger,
	}
}

// Handle list arrived guest operation
func (l ListArrivedGuestHandler) Handle(ctx context.Context) ([]guest.Guest, error) {
	guests, err := l.guestRepo.ListArrivedGuests(ctx)
	l.logger.Debugf("List arrived guest %v", guests)
	return guests, err
}
