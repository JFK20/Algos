package sorting

import (
	"Algos/constants"
	"fmt"
	"sync"
)

func MergeSort(a []int, b []int) []int {
	lengthA, lengthB := len(a), len(b)

	if lengthA == 1 && lengthB == 1 {
		return sort(a[0], b[0])
	} else if lengthA == 1 && lengthB == 2 {
		sorted := sort(b[0], b[1])
		return combine(a, sorted)
	} else if lengthA == 2 && lengthB == 1 {
		sorted := sort(a[0], a[1])
		return combine(sorted, b)
	} else {
		leftMiddle := lengthA / 2
		rightMiddle := lengthB / 2

		left := MergeSort(a[:leftMiddle], a[leftMiddle:])
		right := MergeSort(b[:rightMiddle], b[rightMiddle:])

		return combine(left, right)
	}
}

func sort(a int, b int) []int {
	if constants.AlgoVerbose {
		fmt.Println("start sorting")
	}
	res := make([]int, 2)
	if a > b {
		res[0] = b
		res[1] = a
	} else {
		res[0] = a
		res[1] = b
	}
	if constants.AlgoVerbose {
		fmt.Println(res)
	}
	return res
}

func combine(left []int, right []int) []int {
	ll, lr := len(left), len(right)
	res := make([]int, ll+lr)

	if constants.AlgoVerbose {
		fmt.Println("start merging")
		fmt.Println(left)
		fmt.Println(right)
	}
	ri, li, ai := 0, 0, 0
	for li < ll && ri < lr {
		if left[li] < right[ri] {
			res[ai] = left[li]
			li++
			ai++
		} else {
			res[ai] = right[ri]
			ri++
			ai++
		}
	}

	if li == ll {
		for ; ri < lr; ri++ {
			res[ai] = right[ri]
			ai++
		}
	} else {
		for ; li < ll; li++ {
			res[ai] = left[li]
			ai++
		}
	}
	if constants.AlgoVerbose {
		fmt.Println(res)
	}
	return res
}

func MergeSortAsync(a []int, b []int) []int {
	lengthA, lengthB := len(a), len(b)

	if lengthA == 1 && lengthB == 1 {
		return sort(a[0], b[0])
	} else if lengthA == 1 && lengthB == 2 {
		sorted := sort(b[0], b[1])
		return combine(a, sorted)
	} else if lengthA == 2 && lengthB == 1 {
		sorted := sort(a[0], a[1])
		return combine(sorted, b)
	} else {
		leftMiddle := lengthA / 2
		rightMiddle := lengthB / 2

		var left []int
		var right []int
		//arbitrary number
		if lengthA > 128 && lengthB > 128 {
			wg := sync.WaitGroup{}
			wg.Add(2)

			go func() {
				left = MergeSort(a[:leftMiddle], a[leftMiddle:])
				wg.Done()
			}()

			go func() {
				right = MergeSort(b[:rightMiddle], b[rightMiddle:])
				wg.Done()
			}()

			wg.Wait()
		} else {
			left = MergeSort(a[:leftMiddle], a[leftMiddle:])
			right = MergeSort(b[:rightMiddle], b[rightMiddle:])
		}

		return combine(left, right)
	}
}
