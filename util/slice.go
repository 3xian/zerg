package util

func CompareIntSlice(a, b []int) int {
	n := MinInt(len(a), len(b))
	for i := 0; i < n; i++ {
		if a[i] != b[i] {
			return a[i] - b[i]
		}
	}
	return len(a) - len(b)
}
