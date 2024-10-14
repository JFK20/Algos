package tests

import (
	"Algos/constants"
	"Algos/sorting"
	"fmt"
	"math/rand"
	"slices"
	"testing"
)

func TestQuick(t *testing.T) {

	unsorted := make([]int, 973)

	for i := range unsorted {
		unsorted[i] = rand.Intn(10000)
	}

	if constants.TestVerbose {
		fmt.Printf("unsorted: %v\n", unsorted)
	}

	sorting.QuickSort(unsorted, 0, len(unsorted)-1)

	if !slices.IsSorted(unsorted) {
		t.Errorf("slice is not sorted correctly")
	}

}

func TestQuickAsync(t *testing.T) {

	unsorted := make([]int, 973)

	for i := range unsorted {
		unsorted[i] = rand.Intn(10000)
	}

	if constants.TestVerbose {
		fmt.Printf("unsorted: %v\n", unsorted)
	}

	sorting.QuickSortAsync(unsorted, 0, len(unsorted)-1)

	if !slices.IsSorted(unsorted) {
		t.Errorf("slice is not sorted correctly")
	}

}
