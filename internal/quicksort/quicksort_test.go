package quicksort

import (
	"reflect"
	"testing"
)

var cmp Comparator[int] = func(a, b int) int {
	if a < b {
		return 1
	}
	if a == b {
		return 0
	}
	return -1
}

func TestPartition(t *testing.T) {
	var cmp Comparator[int] = func(a, b int) int {
		if a < b {
			return 1
		}
		if a == b {
			return 0
		}
		return -1
	}
	testCases := []struct {
		slice  []int
		L      int
		R      int
		expect []int
		pivot  int
	}{
		{[]int{}, 0, 0, []int{}, 0},
		{[]int{0}, 0, 0, []int{0}, 0},
		{[]int{1, 2}, 0, 1, []int{1, 2}, 1},
		{[]int{1, 3, 2}, 0, 2, []int{1, 2, 3}, 2},
		{[]int{3, 2, 1}, 0, 2, []int{1, 2, 3}, 1},
		{[]int{5, 4, 3, 2}, 0, 3, []int{2, 3, 4, 5}, 2},
		{[]int{5, 4, 3, 2}, 0, 2, []int{3, 4, 5, 2}, 1},
	}

	for _, testCase := range testCases {
		copied := make([]int, 0, len(testCase.slice))
		copied = append(copied, testCase.slice...)
		pivot := partition(copied, cmp, testCase.L, testCase.R)
		if !reflect.DeepEqual(copied, testCase.expect) || pivot != testCase.pivot {
			t.Errorf("got: (%v, %d)\nexpected: (%v, %d)\ntestCase: %v\n",
				copied, pivot, testCase.expect, testCase.pivot, testCase)
		}
	}
}

func TestQuicksort(t *testing.T) {
	testCases := []struct {
		slice  []int
		expect []int
	}{
		{[]int{}, []int{}},
		{[]int{0}, []int{0}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{1, 3, 2}, []int{1, 2, 3}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{5, 4, 3, 2}, []int{2, 3, 4, 5}},
		{[]int{5, 1, 3, 2}, []int{1, 2, 3, 5}},
	}

	for _, testCase := range testCases {
		copied := make([]int, 0, len(testCase.slice))
		copied = append(copied, testCase.slice...)
		Quicksort(copied, cmp)
		if !reflect.DeepEqual(copied, testCase.expect) {
			t.Errorf("\ngot: %v\nexpected: %v\ntestCase: %v\n",
				copied, testCase.expect, testCase)
		}
	}
}
