package guest

import "context"

// Repository that expected to be implemented by infrastructure layer for manipulating guest in persistent level
type Repository interface {
	// Add guest in persistent level (memory storage, database, etc)
	AddGuest(ctx context.Context, guest Guest) error
	// List guests in persistent level
	ListGuests(ctx context.Context) ([]Guest, error)
	// List arrived guests in persistent level
	ListArrivedGuests(ctx context.Context) ([]Guest, error)
	// Update guest function to store the updated guest in persistent level
	UpdateGuest(ctx context.Context, name string, updateFn func(g *Guest) (*Guest, error)) error
}
