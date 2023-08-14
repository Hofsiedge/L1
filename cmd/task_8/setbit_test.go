package main

import (
	"errors"
	"testing"
)

func TestSetBit(t *testing.T) {
	testCases := []struct {
		value    int64
		bit      int
		one      bool
		expected int64
		err      error
	}{
		{0, 0, true, 1, nil},
		{8, 3, false, 0, nil},
		{8, 2, false, 8, nil},
		{8, 2, true, 12, nil},
		{8, 64, true, 0, ErrInvalidBit},
		{8, -20, true, 0, ErrInvalidBit},
	}
	for _, testCase := range testCases {
		got, err := setBit(testCase.value, testCase.bit, testCase.one)
		// handling test cases for errors
		if testCase.err != nil {
			if errors.Is(err, testCase.err) {
				continue
			}
			t.Errorf(`expected "%v" error, got "%v"`, testCase.err, err)
			continue
		}
		// handling test cases without errors
		if err != nil {
			t.Errorf(`unexpected error: "%v". test case: %v`, err, testCase)
		}
		if got != testCase.expected {
			t.Errorf("expected %d, got %d. test case: %v",
				testCase.expected, got, testCase)
		}
	}
}
