package table

import "testing"

func TestNewTable(t *testing.T) {
	type testcase struct {
		testcase       string
		id             int
		capacity       int
		availableSpace int
		availableSeat  int
		expectedResult Table
		expectedError  error
	}

	testcases := []testcase{
		{
			testcase:       "Normal",
			id:             0,
			capacity:       1,
			availableSpace: 1,
			availableSeat:  1,
			expectedResult: Table{
				id:             0,
				capacity:       1,
				availableSpace: 1,
				availableSeat:  1,
			},
		},
		{
			testcase:       "Negative Capacity",
			id:             0,
			capacity:       -1,
			availableSpace: -1,
			availableSeat:  -1,
			expectedError:  ErrNegativeCapacity,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.testcase, func(t *testing.T) {
			table, err := NewTable(tc.id, tc.capacity, tc.availableSpace, tc.availableSeat)
			if err != tc.expectedError {
				t.Fatalf("expected %v, but got %v", tc.expectedError, err)
			}
			if table.id != tc.expectedResult.id ||
				table.capacity != tc.expectedResult.capacity ||
				table.availableSpace != tc.expectedResult.availableSpace ||
				table.availableSeat != tc.expectedResult.availableSeat {
				t.Errorf("expected %v, but got %v", tc.expectedResult, table)
			}
		})
	}
}

func TestId(t *testing.T) {
	type testcase struct {
		testcase       string
		id             int
		expectedResult int
	}

	testcases := []testcase{
		{
			testcase:       "Normal",
			id:             1,
			expectedResult: 1,
		},
		{
			testcase:       "Normal",
			id:             2,
			expectedResult: 2,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.testcase, func(t *testing.T) {
			g := Table{
				id: tc.id,
			}
			if g.Id() != tc.expectedResult {
				t.Errorf("expected %v, but got %v", tc.expectedResult, g.Id())
			}
		})
	}
}

func TestSetId(t *testing.T) {
	type testcase struct {
		testcase       string
		id             int
		expectedResult int
	}

	testcases := []testcase{
		{
			testcase:       "Normal",
			id:             1,
			expectedResult: 1,
		},
		{
			testcase:       "Normal",
			id:             2,
			expectedResult: 2,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.testcase, func(t *testing.T) {
			g := Table{
				id: 0,
			}
			g.SetId(tc.id)
			if g.id != tc.expectedResult {
				t.Errorf("expected %v, but got %v", tc.expectedResult, g.id)
			}
		})
	}
}

func TestCapacity(t *testing.T) {
	type testcase struct {
		testcase       string
		capacity       int
		expectedResult int
	}

	testcases := []testcase{
		{
			testcase:       "Normal",
			capacity:       1,
			expectedResult: 1,
		},
		{
			testcase:       "Normal",
			capacity:       2,
			expectedResult: 2,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.testcase, func(t *testing.T) {
			g := Table{
				capacity: tc.capacity,
			}
			if g.Capacity() != tc.expectedResult {
				t.Errorf("expected %v, but got %v", tc.expectedResult, g.Capacity())
			}
		})
	}
}
func TestAvailableSpace(t *testing.T) {
	type testcase struct {
		testcase       string
		availableSpace int
		expectedResult int
	}

	testcases := []testcase{
		{
			testcase:       "Normal",
			availableSpace: 1,
			expectedResult: 1,
		},
		{
			testcase:       "Normal",
			availableSpace: 2,
			expectedResult: 2,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.testcase, func(t *testing.T) {
			g := Table{
				availableSpace: tc.availableSpace,
			}
			if g.AvailableSpace() != tc.expectedResult {
				t.Errorf("expected %v, but got %v", tc.expectedResult, g.AvailableSpace())
			}
		})
	}
}

func TestAvailableSeat(t *testing.T) {
	type testcase struct {
		testcase       string
		availableSeat  int
		expectedResult int
	}

	testcases := []testcase{
		{
			testcase:       "Normal",
			availableSeat:  1,
			expectedResult: 1,
		},
		{
			testcase:       "Normal",
			availableSeat:  2,
			expectedResult: 2,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.testcase, func(t *testing.T) {
			g := Table{
				availableSeat: tc.availableSeat,
			}
			if g.AvailableSeat() != tc.expectedResult {
				t.Errorf("expected %v, but got %v", tc.expectedResult, g.AvailableSeat())
			}
		})
	}
}

func TestAllocateSeat(t *testing.T) {
	type testcase struct {
		testcase       string
		table          Table
		allocateNumber int
		remainSpace    int
		expectedError  error
	}

	testcases := []testcase{
		{
			testcase: "Normal",
			table: Table{
				capacity:       5,
				availableSpace: 5,
			},
			allocateNumber: 1,
			remainSpace:    4,
			expectedError:  nil,
		},
		{
			testcase: "Not enough seat",
			table: Table{
				capacity:       5,
				availableSpace: 5,
			},
			remainSpace:    5,
			allocateNumber: 10,
			expectedError:  ErrNotEnoughSeat,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.testcase, func(t *testing.T) {
			err := tc.table.AllocateSeat(tc.allocateNumber)
			if err != tc.expectedError {
				t.Fatalf("expected %v, but got %v", tc.expectedError, err)
			}
			if tc.remainSpace != tc.table.availableSpace {
				t.Errorf("expected %v, but got %v", tc.remainSpace, tc.table.availableSpace)
			}
		})
	}
}

func TestCheckin(t *testing.T) {
	type testcase struct {
		testcase      string
		table         Table
		checkInNumber int
		remainSeat    int
		expectedError error
	}

	testcases := []testcase{
		{
			testcase: "Normal",
			table: Table{
				capacity:       5,
				availableSpace: 5,
				availableSeat:  5,
			},
			checkInNumber: 1,
			remainSeat:    4,
			expectedError: nil,
		},
		{
			testcase: "Not enough seat",
			table: Table{
				capacity:       5,
				availableSpace: 5,
				availableSeat:  5,
			},
			checkInNumber: 10,
			remainSeat:    5,
			expectedError: ErrNotEnoughSeat,
		},
		{
			testcase: "Not enough seat",
			table: Table{
				capacity:       5,
				availableSpace: 5,
				availableSeat:  0,
			},
			checkInNumber: 1,
			remainSeat:    0,
			expectedError: ErrNotEnoughSeat,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.testcase, func(t *testing.T) {
			err := tc.table.Checkin(tc.checkInNumber)
			if err != tc.expectedError {
				t.Fatalf("expected %v, but got %v", tc.expectedError, err)
			}
			if tc.remainSeat != tc.table.availableSeat {
				t.Errorf("expected %v, but got %v", tc.remainSeat, tc.table.availableSeat)
			}
		})
	}
}

func TestCheckOut(t *testing.T) {
	type testcase struct {
		testcase       string
		table          Table
		checkOutNumber int
		remainSeat     int
	}

	testcases := []testcase{
		{
			testcase: "Normal",
			table: Table{
				capacity:       5,
				availableSpace: 5,
				availableSeat:  4,
			},
			checkOutNumber: 1,
			remainSeat:     5,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.testcase, func(t *testing.T) {
			tc.table.Checkout(tc.checkOutNumber)
			if tc.remainSeat != tc.table.availableSeat {
				t.Errorf("expected %v, but got %v", tc.remainSeat, tc.table.availableSeat)
			}
		})
	}
}
