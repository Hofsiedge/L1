package bisect

import "testing"

func TestBisect(t *testing.T) {
	testCases := []struct {
		input    []int
		value    int
		present  bool
		position int
	}{
		{[]int{-10, -1, 2, 3, 5}, 3, true, 3},
		{[]int{-10, -1, 2, 4, 5}, 3, false, 3},
		{[]int{-10}, -100, false, 0},
		{[]int{-10}, 0, false, 1},
		// Bisect returns mid point if comparator(input[M], value) == OrderedEqual
		{[]int{0, 0, 0, 0}, 0, true, 1},
	}

	var comparator Comparator[int] = func(a, b int) ComparatorOrder {
		if a < b {
			return OrderedRight
		}
		if a == b {
			return OrderedEqual
		}
		return OrderedWrong
	}

	for _, testCase := range testCases {
		position, found := Bisect(testCase.input, comparator, testCase.value)
		if position != testCase.position || found != testCase.present {
			t.Errorf("\nexpected: (%d, %v), got: (%d, %v)\ntestCase: %v",
				testCase.position, testCase.present, position, found, testCase)
		}
	}
}
