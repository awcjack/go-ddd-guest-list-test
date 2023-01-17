package query

import (
	"context"

	"github.com/awcjack/getground-backend-assignment/domain/table"
)

// Handler for counting empty seat
type CountEmptySeatHandler struct {
	tableRepo table.Repository
	logger    logger
}

// CountEmptySeatHandler constructor
func NewCountEmptySeatHandler(tableRepo table.Repository, logger logger) *CountEmptySeatHandler {
	return &CountEmptySeatHandler{
		tableRepo: tableRepo,
		logger:    logger,
	}
}

// Handle count empty seat operation
func (c CountEmptySeatHandler) Handle(ctx context.Context) int {
	tables, _ := c.tableRepo.ListTable(ctx)
	count := 0

	c.logger.Debugf("Count Empty Seat tables %v", tables)
	for _, table := range tables {
		count += table.AvailableSeat()
	}

	return count
}
