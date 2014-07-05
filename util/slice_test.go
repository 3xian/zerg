package util

import "testing"

func TestCompareIntSlice(t *testing.T) {
	testCases := []struct {
		a, b []int
		cmp  int
	}{
		{[]int{1, 2}, []int{1, 2}, 0},
		{[]int{2, 2}, []int{1, 2}, 1},
		{[]int{2, 2}, []int{3, 2}, -1},
		{[]int{2, 2}, []int{2, 3}, -1},
		{[]int{1}, []int{1, 2}, -1},
		{[]int{1, 2, 3}, []int{1, 2}, 1},
	}

	for _, tc := range testCases {
		if cmp := CompareIntSlice(tc.a, tc.b); cmp != tc.cmp {
			t.Errorf("CompareIntSlice(%v, %v) got %d, want %d\n",
				tc.a, tc.b, cmp, tc.cmp)
		}
	}
}
