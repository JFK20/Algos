package tests

import (
	"Algos/constants"
	"Algos/sorting"
	"fmt"
	"math/rand"
	"slices"
	"testing"
)

func TestMerge(T *testing.T) {
	unsorted := make([]int, 32)

	for i := range unsorted {
		unsorted[i] = rand.Intn(1000)
	}

	if constants.TestVerbose {
		fmt.Printf("unsorted: %v\n", unsorted)
	}

	got := sorting.MergeSort(unsorted[:len(unsorted)/2], unsorted[len(unsorted)/2:])
	if !slices.IsSorted(got) {
		T.Errorf("slice is not sorted correctly")
	}
}
