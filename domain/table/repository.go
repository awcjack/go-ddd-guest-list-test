package table

import "context"

type Repository interface {
	// Add table to persistent level
	AddTable(ctx context.Context, table *Table) error
	// Get table from persistent level
	GetTable(ctx context.Context, id int) (*Table, error)
	// List all table from persistent level
	ListTable(ctx context.Context) ([]Table, error)
}
