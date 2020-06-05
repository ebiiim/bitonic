// Package bitonic implements a bitonic sorter.
package bitonic

import (
	"sync"
)

// SortOrder represents sort order.
type SortOrder bool

const (
	// Ascending represents ascending order.
	Ascending SortOrder = true
	// Descending represents descending order.
	Descending SortOrder = false
)

// SortInts sorts `x` by `ord` order (concurrent).
// `len(x)` must be a power of 2.
func SortInts(x []int, ord SortOrder) {
	bitonicSort(x, ord, len(x))
}

// SortInts1 sorts `x` by `ord` order (non-concurrent).
// `len(x)` must be a power of 2.
func SortInts1(x []int, ord SortOrder) {
	bitonicSort1(x, ord, len(x))
}

// Threshold is used to decide whether to run concurrently.
const Threshold = 1 << 14

func bitonicSort(x []int, ord SortOrder, ln int) {
	if ln <= 1 {
		return
	}
	mid := ln >> 1
	if mid >= Threshold {
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			bitonicSort(x[:mid], true, mid)
			wg.Done()
		}()
		go func() {
			bitonicSort(x[mid:], false, mid)
			wg.Done()
		}()
		wg.Wait()
	} else {
		bitonicSort(x[:mid], true, mid)
		bitonicSort(x[mid:], false, mid)
	}
	bitonicMerge(x, ord, ln)
}

func bitonicMerge(x []int, ord SortOrder, ln int) {
	if ln <= 1 {
		return
	}
	compareAndSwap(x, ord, ln)
	mid := ln >> 1
	if mid >= Threshold {
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			bitonicMerge(x[:mid], ord, mid)
			wg.Done()
		}()
		go func() {
			bitonicMerge(x[mid:], ord, mid)
			wg.Done()
		}()
		wg.Wait()
	} else {
		bitonicMerge(x[:mid], ord, mid)
		bitonicMerge(x[mid:], ord, mid)
	}
}

func bitonicSort1(x []int, ord SortOrder, ln int) {
	if ln <= 1 {
		return
	}
	mid := ln >> 1
	bitonicSort1(x[:mid], true, mid)
	bitonicSort1(x[mid:], false, mid)
	bitonicMerge1(x, ord, ln)
}

func bitonicMerge1(x []int, ord SortOrder, ln int) {
	if ln <= 1 {
		return
	}
	compareAndSwap(x, ord, ln)
	mid := ln >> 1
	bitonicMerge1(x[:mid], ord, mid)
	bitonicMerge1(x[mid:], ord, mid)
}

func compareAndSwap(x []int, ord SortOrder, ln int) {
	mid := ln >> 1
	for i := 0; i < mid; i++ {
		peer := mid ^ i
		if (x[i] > x[peer]) == ord {
			x[i] = x[i] ^ x[peer]
			x[peer] = x[i] ^ x[peer]
			x[i] = x[i] ^ x[peer]
		}
	}
}
