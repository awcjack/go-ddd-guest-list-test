//go:build integration
// +build integration

// Integration test

package main_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/awcjack/getground-backend-assignment/infrastructure/datastore/testdb"
	"github.com/stretchr/testify/require"
)

// POST /guest_list/{name}
func TestAddGuest(t *testing.T) {
	// Make sure clean the database after each of the test to prevent affect other test
	defer func() {
		if err := testdb.CleanGuests(db); err != nil {
			t.Errorf("error truncating test database tables: %v", err)
		}
		if err := testdb.CleanTables(db); err != nil {
			t.Errorf("error truncating test database tables: %v", err)
		}
	}()

	// Seed default data to database for testing
	err := testdb.SeedTables(db)
	if err != nil {
		t.Fatal("Seed db error", err)
	}

	type testcase struct {
		testcase     string
		pathname     string
		requestBody  string
		expectedBody string
		expectedCode int
	}

	testcases := []testcase{
		{
			testcase:     "Insert 1",
			pathname:     "1",
			requestBody:  `{"table":1,"accompanying_guests":0}`,
			expectedBody: `{"name":"1"}`,
			expectedCode: http.StatusOK,
		},
		{
			testcase:     "Insert 2",
			pathname:     "2",
			requestBody:  `{"table":2,"accompanying_guests":5}`,
			expectedBody: `{"name":"2"}`,
			expectedCode: http.StatusOK,
		},
		{
			testcase:     "Insert 1 Duplicated Guest",
			pathname:     "1",
			requestBody:  `{"table":2,"accompanying_guests":0}`,
			expectedBody: `{"error":"guest already exist"}`,
			expectedCode: http.StatusBadRequest,
		},
		{
			testcase:     "Insert 3 not enough seat",
			pathname:     "3",
			requestBody:  `{"table":1,"accompanying_guests":1}`,
			expectedBody: `{"error":"not enough seat"}`,
			expectedCode: http.StatusBadRequest,
		},
		{
			testcase:     "Insert 3 table doesn't exist",
			pathname:     "3",
			requestBody:  `{"table":10,"accompanying_guests":1}`,
			expectedBody: `{"error":"Table not found from db"}`,
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.testcase, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodPost, "/guest_list/"+tc.pathname, strings.NewReader(tc.requestBody))
			if err != nil {
				t.Errorf("error creating request: %v", err)
			}
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			rootRouter.ServeHTTP(w, req)

			res := w.Result()
			if res.StatusCode != tc.expectedCode {
				t.Errorf("Expect %s status code is %d but received %d", tc.testcase, tc.expectedCode, res.StatusCode)
			}

			defer res.Body.Close()
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Errorf("expected error to be nil got %v", err)
			}
			require.JSONEq(t, tc.expectedBody, string(body))
		})
	}
}

// GET /guest_list
func TestListGuest(t *testing.T) {
	defer func() {
		if err := testdb.CleanGuests(db); err != nil {
			t.Errorf("error truncating test database tables: %v", err)
		}
		if err := testdb.CleanTables(db); err != nil {
			t.Errorf("error truncating test database tables: %v", err)
		}
	}()

	err := testdb.SeedTables(db)
	if err != nil {
		t.Fatal("Seed db error", err)
	}

	err = testdb.SeedGuests(db)
	if err != nil {
		t.Fatal("Seed db error", err)
	}

	type testcase struct {
		testcase     string
		requestBody  string
		expectedBody string
		expectedCode int
	}

	testcases := []testcase{
		{
			testcase:    "Guest list",
			requestBody: ``,
			expectedBody: `{
				"guests":[
					{
						"accompanying_guests":0,
						"name":"John",
						"table":1
					},{
						"accompanying_guests":2,
						"name":"John1",
						"table":2
					},{
						"accompanying_guests":4,
						"name":"John2",
						"table":3
					},{
						"accompanying_guests":0,
						"name":"John3",
						"table":2
					}
				]
			}`,
			expectedCode: http.StatusOK,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.testcase, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, "/guest_list", nil)
			if err != nil {
				t.Errorf("error creating request: %v", err)
			}

			w := httptest.NewRecorder()
			rootRouter.ServeHTTP(w, req)

			res := w.Result()
			if res.StatusCode != tc.expectedCode {
				t.Errorf("Expect %s status code is %d but received %d", tc.testcase, tc.expectedCode, res.StatusCode)
			}

			defer res.Body.Close()
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Errorf("expected error to be nil got %v", err)
			}
			require.JSONEq(t, tc.expectedBody, string(body))
		})
	}
}

// GET /guests
func TestListArrivedGuest(t *testing.T) {
	defer func() {
		if err := testdb.CleanGuests(db); err != nil {
			t.Errorf("error truncating test database tables: %v", err)
		}
		if err := testdb.CleanTables(db); err != nil {
			t.Errorf("error truncating test database tables: %v", err)
		}
	}()

	err := testdb.SeedTables(db)
	if err != nil {
		t.Fatal("Seed db error", err)
	}

	err = testdb.SeedGuests(db)
	if err != nil {
		t.Fatal("Seed db error", err)
	}

	type testcase struct {
		testcase     string
		requestBody  string
		expectedBody string
		expectedCode int
	}

	testcases := []testcase{
		{
			testcase:    "Arrived Guest list",
			requestBody: ``,
			expectedBody: `{
				"guests":[
					{
						"accompanying_guests":3,
						"name":"John1",
						"time_arrived":"2009-11-10 23:00:00"
					},{
						"accompanying_guests":4,
						"name":"John2",
						"time_arrived":"2009-11-11 23:00:00"
					}
				]
			}`,
			expectedCode: http.StatusOK,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.testcase, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, "/guests", nil)
			if err != nil {
				t.Errorf("error creating request: %v", err)
			}

			w := httptest.NewRecorder()
			rootRouter.ServeHTTP(w, req)

			res := w.Result()
			if res.StatusCode != tc.expectedCode {
				t.Errorf("Expect %s status code is %d but received %d", tc.testcase, tc.expectedCode, res.StatusCode)
			}

			defer res.Body.Close()
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Errorf("expected error to be nil got %v", err)
			}
			require.JSONEq(t, tc.expectedBody, string(body))
		})
	}
}

// PUT /guests/{name}
func TestCheckInGuest(t *testing.T) {
	defer func() {
		if err := testdb.CleanGuests(db); err != nil {
			t.Errorf("error truncating test database tables: %v", err)
		}
		if err := testdb.CleanTables(db); err != nil {
			t.Errorf("error truncating test database tables: %v", err)
		}
	}()

	err := testdb.SeedTables(db)
	if err != nil {
		t.Fatal("Seed db error", err)
	}

	err = testdb.SeedGuests(db)
	if err != nil {
		t.Fatal("Seed db error", err)
	}

	type testcase struct {
		testcase     string
		pathname     string
		requestBody  string
		expectedBody string
		expectedCode int
	}

	testcases := []testcase{
		{
			testcase:     "CheckIn John",
			pathname:     "John",
			requestBody:  `{"accompanying_guests":0}`,
			expectedBody: `{"name":"John"}`,
			expectedCode: http.StatusOK,
		},
		{
			testcase:     "CheckIn John1 already Checkin",
			pathname:     "John1",
			requestBody:  `{"accompanying_guests":0}`,
			expectedBody: `{"error":"guest already arrived"}`,
			expectedCode: http.StatusBadRequest,
		},
		{
			testcase:     "CheckIn 1",
			pathname:     "1",
			requestBody:  `{"table":2,"accompanying_guests":0}`,
			expectedBody: `{"error":"guest not exist"}`,
			expectedCode: http.StatusBadRequest,
		},
		{
			testcase:     "CheckIn John3 not enough seat",
			pathname:     "John3",
			requestBody:  `{"accompanying_guests":10}`,
			expectedBody: `{"error":"not enough seat"}`,
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.testcase, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodPut, "/guests/"+tc.pathname, strings.NewReader(tc.requestBody))
			if err != nil {
				t.Errorf("error creating request: %v", err)
			}
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			rootRouter.ServeHTTP(w, req)

			res := w.Result()
			if res.StatusCode != tc.expectedCode {
				t.Errorf("Expect %s status code is %d but received %d", tc.testcase, tc.expectedCode, res.StatusCode)
			}

			defer res.Body.Close()
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Errorf("expected error to be nil got %v", err)
			}
			require.JSONEq(t, tc.expectedBody, string(body))
		})
	}
}

// PUT /guests/{name}
func TestCheckOutGuest(t *testing.T) {
	defer func() {
		if err := testdb.CleanGuests(db); err != nil {
			t.Errorf("error truncating test database tables: %v", err)
		}
		if err := testdb.CleanTables(db); err != nil {
			t.Errorf("error truncating test database tables: %v", err)
		}
	}()

	err := testdb.SeedTables(db)
	if err != nil {
		t.Fatal("Seed db error", err)
	}

	err = testdb.SeedGuests(db)
	if err != nil {
		t.Fatal("Seed db error", err)
	}

	type testcase struct {
		testcase     string
		pathname     string
		requestBody  string
		expectedBody string
		expectedCode int
	}

	testcases := []testcase{
		{
			testcase:     "CheckOut John1",
			pathname:     "John1",
			requestBody:  ``,
			expectedBody: ``,
			expectedCode: http.StatusNoContent,
		},
		{
			testcase:     "CheckOut John haven't Checkin",
			pathname:     "John",
			requestBody:  ``,
			expectedBody: `{"error":"guest doesn't arrived yet"}`,
			expectedCode: http.StatusBadRequest,
		},
		{
			testcase:     "CheckOut 1",
			pathname:     "1",
			requestBody:  ``,
			expectedBody: `{"error":"guest not exist"}`,
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.testcase, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodDelete, "/guests/"+tc.pathname, nil)
			if err != nil {
				t.Errorf("error creating request: %v", err)
			}

			w := httptest.NewRecorder()
			rootRouter.ServeHTTP(w, req)

			res := w.Result()
			if res.StatusCode != tc.expectedCode {
				t.Errorf("Expect %s status code is %d but received %d", tc.testcase, tc.expectedCode, res.StatusCode)
			}

			if tc.expectedCode != http.StatusNoContent {
				defer res.Body.Close()
				body, err := ioutil.ReadAll(res.Body)
				if err != nil {
					t.Errorf("expected error to be nil got %v", err)
				}
				require.JSONEq(t, tc.expectedBody, string(body))
			}
		})
	}
}
