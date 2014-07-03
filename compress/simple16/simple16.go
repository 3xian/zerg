// Package simple16 implements the Simple16 algorithm for sorted integers.
package simple16

import "github.com/3xian/zerg/util"

const (
	bodySize = 28
)

var (
	bodySlots = [][]uint{
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{2, 2, 2, 2, 2, 2, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2},
		{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
		{4, 3, 3, 3, 3, 3, 3, 3, 3},
		{3, 4, 4, 4, 4, 3, 3, 3},
		{4, 4, 4, 4, 4, 4, 4},
		{5, 5, 5, 5, 4, 4},
		{4, 4, 5, 5, 5, 5},
		{6, 6, 6, 5, 5},
		{5, 5, 6, 6, 6},
		{7, 7, 7, 7},
		{10, 9, 9},
		{14, 14},
		{28},
	}
)

func Compress(src []int32, srcOffset int) (dst int32, dstNum int) {
TypeLoop:
	for i, slots := range bodySlots {
		dst = int32(i) << bodySize
		dstNum = util.MinInt(len(slots), len(src)-srcOffset)

		bits := uint(0)
		for j := 0; j < dstNum; j++ {
			if src[srcOffset+j] >= (1 << slots[j]) {
				continue TypeLoop
			}
			dst |= src[srcOffset+j] << bits
			bits += slots[j]
		}
		return
	}

	return 0, 0
}

func Decompress() {
}
