package command

import (
	"context"

	"github.com/awcjack/getground-backend-assignment/domain/table"
)

// Handler for adding table
type AddTableHandler struct {
	tableRepo table.Repository
	logger    logger
}

// AddTableHandler constructor
func NewAddTableHandler(tableRepo table.Repository, logger logger) *AddTableHandler {
	return &AddTableHandler{
		tableRepo: tableRepo,
		logger:    logger,
	}
}

// Handle Add table operation
func (a *AddTableHandler) Handle(ctx context.Context, capacity int) (table.Table, error) {
	newTable, err := table.NewTable(0, capacity, capacity, capacity)
	a.logger.Debugf("Add Table %v", newTable)
	if err != nil {
		return table.Table{}, err
	}

	err = a.tableRepo.AddTable(ctx, &newTable)
	if err != nil {
		return table.Table{}, err
	}

	return newTable, nil
}
