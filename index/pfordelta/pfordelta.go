package pfordelta

// Compress compresses numbers by PForDelta algorithm.
//func Compress(src []int32) []byte { bitSize := getStrictBitSize(src) }

// Decompress returns numbers that compressed by PForDelta algorithm.
//func Decompress(src []byte) []int32 { }

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
