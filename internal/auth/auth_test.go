package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	testHeader1 := http.Header{}
	testHeader2 := http.Header{}
	testHeader3 := http.Header{}
	testHeader4 := http.Header{}
	testHeader1.Set("", "")
	testHeader2.Set("Authorization", "")
	testHeader3.Set("Authorization", "teset api")
	testHeader4.Set("Authorization", "ApiKey correctkey")
	cases := []struct {
		input    http.Header
		expected struct {
			key string
			err error
		}
	}{
		{input: testHeader1, expected: struct {
			key string
			err error
		}{key: "", err: ErrNoAuthHeaderIncluded}},
		{input: testHeader2, expected: struct {
			key string
			err error
		}{key: "", err: ErrNoAuthHeaderIncluded}},
		{input: testHeader2, expected: struct {
			key string
			err error
		}{key: "", err: ErrNoAuthHeaderIncluded}},
		{input: testHeader4, expected: struct {
			key string
			err error
		}{key: "correctkey", err: nil}},
	}

	for i, testCase := range cases {
		actual, err := GetAPIKey(testCase.input)

		if actual != testCase.expected.key || err != testCase.expected.err {
			t.Errorf("expected key %d: %v, got: %v, expected err: %v, got: %v", i, testCase.expected.key, actual, testCase.expected.err, err)
		}
	}
}
