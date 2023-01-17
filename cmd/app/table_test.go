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

// POST /tables
func TestAddTable(t *testing.T) {
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

	type testcase struct {
		testcase     string
		requestBody  string
		expectedBody string
		expectedCode int
	}

	testcases := []testcase{
		{
			testcase:     "Insert 1",
			requestBody:  `{"capacity":10}`,
			expectedBody: `{"capacity":10,"id":4}`,
			expectedCode: http.StatusOK,
		},
		{
			testcase:     "Insert 2",
			requestBody:  `{"capacity":8}`,
			expectedBody: `{"capacity":8,"id":5}`,
			expectedCode: http.StatusOK,
		},
		{
			testcase:     "Negative capacity",
			requestBody:  `{"capacity":-10}`,
			expectedBody: `{"error":"negative capacity"}`,
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.testcase, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodPost, "/tables", strings.NewReader(tc.requestBody))
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

// GET /seats_empty
func TestGetEmptySeat(t *testing.T) {
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
		t.Fatal("Seed table db error", err)
	}
	err = testdb.SeedGuests(db)
	if err != nil {
		t.Fatal("Seed guest db error", err)
	}

	type testcase struct {
		testcase     string
		requestBody  string
		expectedBody string
		expectedCode int
	}

	testcases := []testcase{
		{
			testcase:     "Get empty seat",
			requestBody:  ``,
			expectedBody: `{"seats_empty":7}`,
			expectedCode: http.StatusOK,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.testcase, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, "/seats_empty", nil)
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
