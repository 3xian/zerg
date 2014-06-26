package pfordelta

import "testing"

func TestGetStrictBitSize(t *testing.T) {
	testCases := []struct {
		numbers []int32
		bitSize int
	}{
		{[]int32{0, 1, 2, 0, 3, 1, 3}, 2},
		{[]int32{0}, 0},
		{[]int32{0, 0, 0}, 0},
		{[]int32{1}, 1},
		{[]int32{1, 0, 1}, 1},
		{[]int32{0, 1, 2, 0, 3, 1, 123}, 7},
		{[]int32{2000000000, 2, 0, 3, 1, 123}, 31},
		{[]int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 2000000000}, 4},
		{[]int32{1, 2, 3, 4, 5, 6, 7, 8, 2000000000, 2000000000}, 31},
	}

	for _, tc := range testCases {
		if got := getStrictBitSize(tc.numbers); got != tc.bitSize {
			t.Errorf("%v got %d, want %d\n", tc.numbers, got, tc.bitSize)
		}
	}
}
