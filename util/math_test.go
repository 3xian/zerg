package util

import "testing"

func TestMinInt(t *testing.T) {
	testCases := []struct {
		ints   []int
		minInt int
	}{
		{[]int{1, 2}, 1},
		{[]int{2, 1}, 1},
		{[]int{10, 10}, 10},
		{[]int{-1, -1}, -1},
		{[]int{-1, -2}, -2},
		{[]int{-2, -1}, -2},
	}

	for _, tc := range testCases {
		if got := MinInt(tc.ints[0], tc.ints[1]); got != tc.minInt {
			t.Errorf("MinInt(%v) got %d, want %d\n", tc.ints, got, tc.minInt)
		}
	}
}
