package simple16

import (
	"github.com/3xian/zerg/util"
	"math/rand"
	"testing"
	"unsafe"
)

func TestCompressOneBlock(t *testing.T) {
	testCases := []struct {
		raw    []int
		cpr    uint32
		cprNum int
	}{
		{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 0x0FFFFFFF, 28},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 0x0FFFFFFF, 28},
		{[]int{0x0FFFFFFE}, 0xFFFFFFFE, 1},
		{[]int{0x0FFFFFFF}, 0xFFFFFFFF, 1},
		{[]int{0x0FFFFFFF + 1}, 0, 0},
		{[]int{1, 15, 2}, 0x60000179, 3},
	}

	for _, tc := range testCases {
		if cpr, cprNum := compressOneBlock(tc.raw); cpr != tc.cpr || cprNum != tc.cprNum {
			t.Errorf("compressOneBlock(%v) got [%b %d], want [%b %d]\n",
				tc.raw, cpr, cprNum, tc.cpr, tc.cprNum)
		}
	}
}

func TestDecompressOneBlock(t *testing.T) {
	testCases := []struct {
		cpr uint32
		raw []int
	}{
		{0x0FFFFFFF, []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}},
		{0xFFFFFFFE, []int{0x0FFFFFFE}},
		{0xFFFFFFFF, []int{0x0FFFFFFF}},
		{0x60000179, []int{1, 15, 2}},
	}

	for _, tc := range testCases {
		if raw := decompressOneBlock(tc.cpr); util.CompareIntSlice(raw, tc.raw) != 0 {
			t.Errorf("decompressOneBlock(%b) got %v, want %v\n",
				tc.cpr, raw, tc.raw)
		}
	}
}

func TestCompressAndDecompress(t *testing.T) {
	r := rand.New(rand.NewSource(99))
	for i := 0; i < 20; i++ {
		n := r.Int31n(500) + 500
		raw := make([]int, n)
		for i := range raw {
			raw[i] = r.Int()%500 + 1
		}

		cpr := Compress(raw)
		t.Logf("%d elemets, from %d bytes to %d bytes\n", len(raw),
			len(raw)*int(unsafe.Sizeof(raw[0])), len(cpr)*int(unsafe.Sizeof(cpr[0])))

		if dcpr := Decompress(cpr); util.CompareIntSlice(raw, dcpr) != 0 {
			t.Errorf("Decompress(%v) got %v, want %v\n", cpr, dcpr, raw)
		}
	}
}
