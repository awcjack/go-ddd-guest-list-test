package command

import (
	"context"

	"github.com/awcjack/getground-backend-assignment/domain/guest"
)

// Handler for adding guest
type AddGuestHandler struct {
	guestRepo guest.Repository
	logger    logger
}

// AddGuestHandler constructor
func NewAddGuestHandler(guestRepo guest.Repository, logger logger) *AddGuestHandler {
	return &AddGuestHandler{
		guestRepo: guestRepo,
		logger:    logger,
	}
}

// Handle add guest operation
func (a AddGuestHandler) Handle(ctx context.Context, guest guest.Guest) error {
	err := a.guestRepo.AddGuest(ctx, guest)
	a.logger.Debugf("Add Guest err %v", err)
	return err
}

// Handler for checkin guest
type CheckInGuestHandler struct {
	guestRepo guest.Repository
	logger    logger
}

// CheckInGuestHandler constructor
func NewCheckInGuestHandler(guestRepo guest.Repository, logger logger) *CheckInGuestHandler {
	return &CheckInGuestHandler{
		guestRepo: guestRepo,
		logger:    logger,
	}
}

// Handle check in guest
func (c CheckInGuestHandler) Handle(ctx context.Context, name string, accompanyingGuest int) error {
	return c.guestRepo.UpdateGuest(ctx, name, func(g *guest.Guest) (*guest.Guest, error) {
		c.logger.Debugf("Checkin Guest Before Update %v", g)
		if err := g.Arrive(accompanyingGuest + 1); err != nil {
			return nil, err
		}
		c.logger.Debugf("Checkin Guest After Update %v", g)
		return g, nil
	})
}

// Handler for checkout guest
type CheckOutGuestHandler struct {
	guestRepo guest.Repository
	logger    logger
}

// CheckOutGuestHandler constructor
func NewCheckOutGuestHandler(guestRepo guest.Repository, logger logger) *CheckOutGuestHandler {
	return &CheckOutGuestHandler{
		guestRepo: guestRepo,
		logger:    logger,
	}
}

// Handle check out guest
func (c CheckOutGuestHandler) Handle(ctx context.Context, name string) error {
	return c.guestRepo.UpdateGuest(ctx, name, func(g *guest.Guest) (*guest.Guest, error) {
		c.logger.Debugf("Checkout Guest Before Update %v", g)
		if err := g.Leave(); err != nil {
			return nil, err
		}
		c.logger.Debugf("Checkout Guest After Update %v", g)
		return g, nil
	})
}
