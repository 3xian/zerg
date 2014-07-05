// Package simple16 implements the Simple16 algorithm.
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

func Compress(raw []int) (cpr []uint32) {
	for i := 0; i < len(raw); {
		block, num := compressOneBlock(raw[i:])
		cpr = append(cpr, block)
		i += num
	}
	return
}

func Decompress(cpr []uint32) (raw []int) {
	for _, block := range cpr {
		for _, elem := range decompressOneBlock(block) {
			raw = append(raw, elem)
		}
	}
	return
}

func compressOneBlock(raw []int) (cpr uint32, cprNum int) {
TypeLoop:
	for i, slots := range bodySlots {
		cpr = uint32(i) << bodySize
		cprNum = util.MinInt(len(slots), len(raw))

		bitShift := uint(0)
		for j := 0; j < cprNum; j++ {
			if raw[j] >= (1 << slots[j]) {
				continue TypeLoop
			}
			cpr |= uint32(raw[j]) << bitShift
			bitShift += slots[j]
		}
		return
	}

	return 0, 0
}

func decompressOneBlock(cpr uint32) (raw []int) {
	slotsType := cpr >> bodySize
	bitShift := uint(0)
	for _, slot := range bodySlots[slotsType] {
		elem := cpr >> bitShift & ((1 << slot) - 1)
		if elem == 0 {
			break
		}
		raw = append(raw, int(elem))
		bitShift += slot
	}
	return
}
