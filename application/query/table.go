package query

import (
	"context"

	"github.com/awcjack/getground-backend-assignment/domain/table"
)

// Handler for getting table
type GetTableHandler struct {
	tableRepo table.Repository
	logger    logger
}

// GetTableHandler constructor
func NewGetTableHandler(tableRepo table.Repository, logger logger) *GetTableHandler {
	return &GetTableHandler{
		tableRepo: tableRepo,
		logger:    logger,
	}
}

// Handle get table operation
func (g GetTableHandler) Handle(ctx context.Context, id int) (*table.Table, error) {
	table, err := g.tableRepo.GetTable(ctx, id)
	g.logger.Debugf("Get Table %v", table)
	return table, err
}
