package bitonic_test

import (
	"math/rand"
	"os"
	"reflect"
	sort "sort"
	"strconv"
	"testing"

	"github.com/ebiiim/bitonic"
)

func makeInts(ln int) (ints []int) {
	rand.Seed(12345)
	ints = make([]int, ln)
	for i := 0; i < ln; i++ {
		v := int(rand.Int31())
		ints[i] = v
	}
	return
}

func makeIntsWithSorted(ln int, ord bitonic.SortOrder) (ints, sortedInts []int) {
	ints = makeInts(ln)
	sortedInts = make([]int, ln)
	_ = copy(sortedInts, ints)
	if ord {
		sort.Ints(sortedInts)
	} else {
		sort.Sort(sort.Reverse(sort.IntSlice(sortedInts)))
	}
	return ints, sortedInts
}

func TestSortInt(t *testing.T) {
	var (
		a1, a2 = makeIntsWithSorted(bitonic.Threshold<<1, bitonic.Ascending)
		d1, d2 = makeIntsWithSorted(bitonic.Threshold<<1, bitonic.Descending)
	)
	cases := []struct {
		name string
		ord  bitonic.SortOrder
		in   []int
		want []int
	}{
		{"zero", bitonic.Ascending, []int{}, []int{}},
		{"asc", bitonic.Ascending, []int{55, 44, 66, 88, 22, 11, 33, 77}, []int{11, 22, 33, 44, 55, 66, 77, 88}},
		{"desc", bitonic.Descending, []int{55, 44, 66, 88, 22, 11, 33, 77}, []int{88, 77, 66, 55, 44, 33, 22, 11}},
		{"asc_long", bitonic.Ascending, a1, a2},
		{"desc_long", bitonic.Descending, d1, d2},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			bitonic.SortInts(c.in, c.ord)
			if !reflect.DeepEqual(c.in, c.want) {
				t.Error("not equal")
			}
		})
	}
}

func TestSortInt1(t *testing.T) {
	var (
		a1, a2 = makeIntsWithSorted(bitonic.Threshold<<1, bitonic.Ascending)
		d1, d2 = makeIntsWithSorted(bitonic.Threshold<<1, bitonic.Descending)
	)
	cases := []struct {
		name string
		ord  bitonic.SortOrder
		in   []int
		want []int
	}{
		{"zero", bitonic.Ascending, []int{}, []int{}},
		{"asc", bitonic.Ascending, []int{55, 44, 66, 88, 22, 11, 33, 77}, []int{11, 22, 33, 44, 55, 66, 77, 88}},
		{"desc", bitonic.Descending, []int{55, 44, 66, 88, 22, 11, 33, 77}, []int{88, 77, 66, 55, 44, 33, 22, 11}},
		{"asc_long", bitonic.Ascending, a1, a2},
		{"desc_long", bitonic.Descending, d1, d2},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			bitonic.SortInts1(c.in, c.ord)
			if !reflect.DeepEqual(c.in, c.want) {
				t.Error("not equal")
			}
		})
	}
}

var benchSize = func() int {
	if b, err := strconv.Atoi(os.Getenv("BITS")); err == nil {
		return b
	}
	return 26 // default value
}()

func BenchmarkSortInt(b *testing.B) {
	b.Logf("bitonic.SortInts %d integers", 1<<benchSize)
	x := makeInts(1<<benchSize)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bitonic.SortInts(x, bitonic.Ascending)
	}
}

func BenchmarkSortInt1(b *testing.B) {
	b.Logf("bitonic.SortInts1 %d integers", 1<<benchSize)
	x := makeInts(1<<benchSize)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bitonic.SortInts1(x, bitonic.Ascending)
	}
}

func BenchmarkGoSort(b *testing.B) {
	b.Logf("sort.Sort %d integers", 1<<benchSize)
	x := makeInts(1<<benchSize)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.Ints(x)
	}
}
