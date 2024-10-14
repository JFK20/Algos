package benchmarks

import (
	"Algos/constants"
	"Algos/sorting"
	"fmt"
	"math/rand"
	"slices"
	"testing"
)

func BenchmarkMerge(b *testing.B) {

	b.StopTimer()

	if b.N == 1 {
		b.N = 2
	}

	unsorted := make([]int, b.N)

	for i := range unsorted {
		unsorted[i] = rand.Intn(1000)
	}

	if constants.TestVerbose {
		fmt.Printf("unsorted: %v\n", unsorted)
	}

	b.StartTimer()
	got := sorting.MergeSort(unsorted[:len(unsorted)/2], unsorted[len(unsorted)/2:])
	b.StopTimer()
	if !slices.IsSorted(got) {
		b.Errorf("slice is not sorted correctly")
	}
}

func BenchmarkMergeAsync(b *testing.B) {
	if b.N == 1 {
		b.N = 2
	}

	unsorted := make([]int, b.N)

	for i := range unsorted {
		unsorted[i] = rand.Intn(1000)
	}

	if constants.TestVerbose {
		fmt.Printf("unsorted: %v\n", unsorted)
	}

	got := sorting.MergeSortAsync(unsorted[:len(unsorted)/2], unsorted[len(unsorted)/2:])
	if !slices.IsSorted(got) {
		b.Errorf("slice is not sorted correctly")
	}
}
