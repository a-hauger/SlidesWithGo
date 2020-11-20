package tdt

import (
	"fmt"
	"testing"
)

func TestWithNoAttachedFunc(t *testing.T) {
	testCases := []struct {
		name string
		arg  string
		want string
	}{
		{"name1", "arg1", "want1"},
	}
	for _, tc := range testCases {
		fmt.Errorf("Name: %v\nArg: %v\nWant: %v", tc.name, tc.arg, tc.want)
	}
}
