package query_test

import (
	"context"
	"testing"
)

func TestCountEmptySeat(t *testing.T) {
	dep := newTableDependencies()
	dep.countEmptySeatHandler.Handle(context.Background())

	if dep.Repository.CalledAddTableTime != 0 ||
		dep.Repository.CalledGetTableTime != 0 ||
		dep.Repository.CalledListTableTime != 1 {
		t.Errorf("Expected only call add table once, but got %v", dep.Repository)
	}
}
