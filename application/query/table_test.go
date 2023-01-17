package query_test

import (
	"context"
	"testing"

	"github.com/awcjack/getground-backend-assignment/application/query"
	"github.com/awcjack/getground-backend-assignment/domain/table"
	"github.com/sirupsen/logrus"
)

func TestGetTable(t *testing.T) {
	dep := newTableDependencies()
	dep.getTableHandler.Handle(context.Background(), 1)

	if dep.Repository.CalledAddTableTime != 0 ||
		dep.Repository.CalledGetTableTime != 1 ||
		dep.Repository.CalledListTableTime != 0 {
		t.Errorf("Expected only call add table once, but got %v", dep.Repository)
	}
}

// struct that keep the mocked repository implementation and handler
type tableDependencies struct {
	Repository            *tableRepoMock
	getTableHandler       *query.GetTableHandler
	countEmptySeatHandler *query.CountEmptySeatHandler
}

// tableDependencies constructor
func newTableDependencies() tableDependencies {
	repository := &tableRepoMock{}
	logger := logrus.NewEntry(logrus.StandardLogger())

	return tableDependencies{
		Repository:            repository,
		getTableHandler:       query.NewGetTableHandler(repository, logger),
		countEmptySeatHandler: query.NewCountEmptySeatHandler(repository, logger),
	}
}

// Mock repository implementation for tracking the handler called correct function
type tableRepoMock struct {
	CalledAddTableTime  int
	CalledGetTableTime  int
	CalledListTableTime int
}

func (t *tableRepoMock) AddTable(ctx context.Context, table *table.Table) error {
	t.CalledAddTableTime++
	return nil
}

func (t *tableRepoMock) GetTable(ctx context.Context, id int) (*table.Table, error) {
	t.CalledGetTableTime++
	return nil, nil
}

func (t *tableRepoMock) ListTable(ctx context.Context) ([]table.Table, error) {
	t.CalledListTableTime++
	return nil, nil
}
