package client_test

import "testing"

type testCase struct {
	name string
}

func Test_Search(t *testing.T) {
	for _, tc := range tcs() {
		t.Run(tc.name, func(t *testing.T) {
		})
	}
}

func tcs() []testCase {
	return []testCase{}
}
