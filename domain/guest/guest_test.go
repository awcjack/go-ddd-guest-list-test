package guest

import (
	"testing"
	"time"

	"github.com/awcjack/getground-backend-assignment/domain/table"
)

func TestName(t *testing.T) {
	type testcase struct {
		testcase       string
		name           string
		expectedResult string
	}

	testcases := []testcase{
		{
			testcase:       "Normal",
			name:           "John",
			expectedResult: "John",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.testcase, func(t *testing.T) {
			g := Guest{
				name:        tc.name,
				guestNumber: 0,
				table:       nil,
			}
			if g.Name() != tc.expectedResult {
				t.Errorf("expected %v, but got %v", tc.expectedResult, g.Name())
			}
		})
	}
}

func TestTable(t *testing.T) {
	type testcase struct {
		testcase       string
		table          *table.Table
		expectedResult *table.Table
	}

	table1, _ := table.NewTable(1, 1, 1, 1)
	table2, _ := table.NewTable(2, 1, 1, 1)
	testcases := []testcase{
		{
			testcase:       "Normal",
			table:          &table1,
			expectedResult: &table1,
		},
		{
			testcase:       "Normal",
			table:          &table2,
			expectedResult: &table2,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.testcase, func(t *testing.T) {
			g := Guest{
				name:        "John",
				guestNumber: 0,
				table:       tc.table,
			}
			if g.Table() != tc.expectedResult {
				t.Errorf("expected %v, but got %v", tc.expectedResult, g.Table())
			}
		})
	}
}

func TestGuestNumber(t *testing.T) {
	type testcase struct {
		testcase       string
		guestNumber    int
		expectedResult int
	}

	testcases := []testcase{
		{
			testcase:       "Normal",
			guestNumber:    1,
			expectedResult: 1,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.testcase, func(t *testing.T) {
			g := Guest{
				name:        "John",
				guestNumber: tc.guestNumber,
				table:       nil,
			}
			if g.GuestNumber() != tc.expectedResult {
				t.Errorf("expected %v, but got %v", tc.expectedResult, g.GuestNumber())
			}
		})
	}
}

func TestArrived(t *testing.T) {
	type testcase struct {
		testcase       string
		arrived        bool
		expectedResult bool
	}

	testcases := []testcase{
		{
			testcase:       "Normal",
			arrived:        true,
			expectedResult: true,
		},
		{
			testcase:       "Normal",
			arrived:        false,
			expectedResult: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.testcase, func(t *testing.T) {
			g := Guest{
				name:        "John",
				guestNumber: 0,
				table:       nil,
				arrived:     tc.arrived,
			}
			if g.Arrived() != tc.expectedResult {
				t.Errorf("expected %v, but got %v", tc.expectedResult, g.Arrived())
			}
		})
	}
}

func TestArrivedNumber(t *testing.T) {
	type testcase struct {
		testcase       string
		arrivedNumber  int
		expectedResult int
	}

	testcases := []testcase{
		{
			testcase:       "Normal",
			arrivedNumber:  1,
			expectedResult: 1,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.testcase, func(t *testing.T) {
			g := Guest{
				name:          "John",
				guestNumber:   0,
				table:         nil,
				arrivedNumber: tc.arrivedNumber,
			}
			if g.ArrivedNumber() != tc.expectedResult {
				t.Errorf("expected %v, but got %v", tc.expectedResult, g.ArrivedNumber())
			}
		})
	}
}

func TestArrivedTime(t *testing.T) {
	type testcase struct {
		testcase       string
		arrivedTime    time.Time
		expectedResult time.Time
	}

	fakeTime, _ := time.Parse("2006-01-02 15:04:05", "2006-01-02 15:04:05")
	testcases := []testcase{
		{
			testcase:       "Normal",
			arrivedTime:    fakeTime,
			expectedResult: fakeTime,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.testcase, func(t *testing.T) {
			g := Guest{
				name:        "John",
				guestNumber: 0,
				table:       nil,
				arrivedTime: tc.arrivedTime,
			}
			if g.ArrivedTime() != tc.expectedResult {
				t.Errorf("expected %v, but got %v", tc.expectedResult, g.ArrivedTime())
			}
		})
	}
}

func TestSetArrived(t *testing.T) {
	type testcase struct {
		testcase       string
		arrived        bool
		expectedResult bool
	}

	testcases := []testcase{
		{
			testcase:       "Normal",
			arrived:        true,
			expectedResult: true,
		},
		{
			testcase:       "Normal",
			arrived:        false,
			expectedResult: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.testcase, func(t *testing.T) {
			g := Guest{
				name:        "John",
				guestNumber: 0,
				table:       nil,
			}
			g.SetArrived(tc.arrived)
			if g.arrived != tc.expectedResult {
				t.Errorf("expected %v, but got %v", tc.expectedResult, g.arrived)
			}
		})
	}
}

func TestSetArriveNumber(t *testing.T) {
	type testcase struct {
		testcase       string
		arrivedNumber  int
		expectedResult int
	}

	testcases := []testcase{
		{
			testcase:       "Normal",
			arrivedNumber:  1,
			expectedResult: 1,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.testcase, func(t *testing.T) {
			g := Guest{
				name:        "John",
				guestNumber: 0,
				table:       nil,
			}
			g.SetArriveNumber(tc.arrivedNumber)
			if g.arrivedNumber != tc.expectedResult {
				t.Errorf("expected %v, but got %v", tc.expectedResult, g.arrivedNumber)
			}
		})
	}
}

func TestSetArrivedTime(t *testing.T) {
	type testcase struct {
		testcase       string
		arrivedTime    time.Time
		expectedResult time.Time
	}

	fakeTime, _ := time.Parse("2006-01-02 15:04:05", "2006-01-02 15:04:05")
	testcases := []testcase{
		{
			testcase:       "Normal",
			arrivedTime:    fakeTime,
			expectedResult: fakeTime,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.testcase, func(t *testing.T) {
			g := Guest{
				name:        "John",
				guestNumber: 0,
				table:       nil,
			}
			g.SetArrivedTime(tc.arrivedTime)
			if g.arrivedTime != tc.expectedResult {
				t.Errorf("expected %v, but got %v", tc.expectedResult, g.arrivedTime)
			}
		})
	}
}

func TestArrive(t *testing.T) {
	type testcase struct {
		testcase       string
		arrived        bool
		guestNumber    int
		table          *table.Table
		expectedResult Guest
		expectedError  error
	}

	table1, _ := table.NewTable(1, 1, 1, 1)
	table2, _ := table.NewTable(2, 1, 1, 1)
	table3, _ := table.NewTable(3, 1, 1, 1)
	testcases := []testcase{
		{
			testcase:    "Normal",
			arrived:     false,
			guestNumber: 1,
			table:       &table1,
			expectedResult: Guest{
				name:          "John",
				guestNumber:   1,
				table:         &table1,
				arrived:       true,
				arrivedNumber: 1,
			},
		},
		{
			testcase:      "Arrved",
			arrived:       true,
			guestNumber:   1,
			table:         &table2,
			expectedError: ErrGuestAlreadyArrived,
			expectedResult: Guest{
				name:          "John",
				guestNumber:   1,
				table:         &table2,
				arrived:       true,
				arrivedNumber: 1,
			},
		},
		{
			testcase:      "Table insufficient seat",
			arrived:       true,
			guestNumber:   2,
			table:         &table3,
			expectedError: table.ErrNotEnoughSeat,
			expectedResult: Guest{
				name:        "John",
				guestNumber: 1,
				table:       &table3,
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.testcase, func(t *testing.T) {
			g := Guest{
				name:        "John",
				guestNumber: 1,
				table:       tc.table,
			}
			err := g.Arrive(tc.guestNumber)
			if err != nil && err != tc.expectedError {
				t.Fatalf("expected error %v, but got %v", tc.expectedError, err)
			}

			if g.name != tc.expectedResult.name ||
				g.guestNumber != tc.expectedResult.guestNumber ||
				g.table != tc.expectedResult.table ||
				g.arrived != tc.expectedResult.arrived ||
				g.arrivedNumber != tc.expectedResult.arrivedNumber {
				t.Errorf("expected %v, but got %v", tc.expectedResult, g)
			}
		})
	}
}

func TestLeave(t *testing.T) {
	type testcase struct {
		testcase       string
		guest          Guest
		expectedResult Guest
		expectedError  error
	}

	table1, _ := table.NewTable(1, 2, 0, 0)
	testcases := []testcase{
		{
			testcase: "Normal",
			guest: Guest{
				name:          "John",
				guestNumber:   1,
				table:         &table1,
				arrived:       true,
				arrivedNumber: 1,
			},
			expectedResult: Guest{
				name:          "John",
				guestNumber:   1,
				table:         &table1,
				arrived:       false,
				arrivedNumber: 0,
			},
		},
		{
			testcase: "Not Arrived",
			guest: Guest{
				name:          "John",
				guestNumber:   1,
				table:         &table1,
				arrived:       false,
				arrivedNumber: 0,
			},
			expectedError: ErrGuestNotArrived,
			expectedResult: Guest{
				name:          "John",
				guestNumber:   1,
				table:         &table1,
				arrived:       false,
				arrivedNumber: 0,
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.testcase, func(t *testing.T) {
			err := tc.guest.Leave()
			if err != nil && err != tc.expectedError {
				t.Fatalf("expected error %v, but got %v", tc.expectedError, err)
			}

			if tc.guest.name != tc.expectedResult.name ||
				tc.guest.guestNumber != tc.expectedResult.guestNumber ||
				tc.guest.table != tc.expectedResult.table ||
				tc.guest.arrived != tc.expectedResult.arrived ||
				tc.guest.arrivedNumber != tc.expectedResult.arrivedNumber {
				t.Errorf("expected %v, but got %v", tc.expectedResult, tc)
			}
		})
	}
}

func TestIsZero(t *testing.T) {
	type testcase struct {
		testcase       string
		name           string
		expectedResult bool
	}

	testcases := []testcase{
		{
			testcase:       "Normal",
			name:           "",
			expectedResult: true,
		},
		{
			testcase:       "Normal",
			name:           "John",
			expectedResult: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.testcase, func(t *testing.T) {
			g := Guest{
				name:        tc.name,
				guestNumber: 0,
				table:       nil,
			}
			if g.IsZero() != tc.expectedResult {
				t.Errorf("expected %v, but got %v", tc.expectedResult, g.IsZero())
			}
		})
	}
}
