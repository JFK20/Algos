package benchmarks

import (
	"Algos/constants"
	"Algos/sorting"
	"fmt"
	"math/rand"
	"slices"
	"testing"
)

func BenchmarkBubble(b *testing.B) {

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
	sorting.BubbleSort(unsorted)
	b.StopTimer()
	if !slices.IsSorted(unsorted) {
		b.Errorf("slice is not sorted correctly")
	}
}
