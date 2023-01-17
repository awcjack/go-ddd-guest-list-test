package config

import (
	"os"
	"testing"
)

func TestGetStringConfigWithDefault(t *testing.T) {
	type testcase struct {
		testcase       string
		key            string
		keyEnv         string
		defaultValue   string
		expectedResult string
	}

	testcases := []testcase{
		{
			testcase:       "Normal",
			key:            "testKey",
			keyEnv:         "test",
			defaultValue:   "testDefault",
			expectedResult: "test",
		},
		{
			testcase:       "Empty env",
			key:            "testKey",
			keyEnv:         "",
			defaultValue:   "testDefault",
			expectedResult: "testDefault",
		},
		{
			testcase:       "Both empty",
			key:            "testKey",
			keyEnv:         "",
			defaultValue:   "",
			expectedResult: "",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.testcase, func(t *testing.T) {
			os.Setenv(tc.key, tc.keyEnv)
			v, err := getStringConfigWithDefault(tc.key, tc.defaultValue)
			if err != nil {
				t.Fatal(err)
			}

			if v != tc.expectedResult {
				t.Errorf("expected %v, but got %v", tc.expectedResult, v)
			}
		})
	}
}
