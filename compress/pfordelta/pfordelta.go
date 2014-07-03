// Package pfordelta implements the optimized PForDelta algorithm.
package pfordelta

const (
	blockSize = 128
)

var (
	possibleBitSize = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 16, 20, 28}
)

// Compress compresses integers using PForDelta.
func Compress(src []int32) []byte {
	bestBitSize := -1
	bestTotalSize := -1

	for _, bitSize := range possibleBitSize {
		if totalSize := estimateCompressedSize(src, bitSize); bestTotalSize < 0 || totalSize < bestTotalSize {
			bestBitSize = bitSize
			bestTotalSize = totalSize
		}
	}

	return CompressWithBitSize(src, bestBitSize)
}

// CompressWithBitSize compresses integers using PForDelta with specific bit size.
func CompressWithBitSize(src []int32, bitSize int) []byte {
}

// Decompress decompresses integers using PForDelta.
//func Decompress(src []byte) []int32 { }

func estimateCompressedSize(src []int32, bitSize int) int {
	return 0 // TODO
}

// getStrictBitSize calculates strict bit size for 90% small numbers.
func getStrictBitSize(src []int32) int {
	counter := [32]int{}
	for _, x := range src {
		size := 0
		for i := 30; i >= 0; i-- {
			if x&(1<<uint(i)) != 0 {
				size = i + 1
				break
			}
		}
		counter[size] += 1
	}

	sum := 0
	for i, c := range counter {
		sum += c

		// Check if sum/len >= 0.9.
		if sum*10 >= len(src)*9 {
			return i
		}
	}

	// Impossible to arrive here.
	return 31
}
