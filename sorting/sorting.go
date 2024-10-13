package sorting

import (
	"Algos/Constants"
	"errors"
	"fmt"
)

func MergeSort(a []int, b []int) []int {
	lengthA, lengthB := len(a), len(b)
	if lengthA != lengthB {
		errors.New(("Arrays lengths are not equal"))
	}
	if lengthA == 1 {
		if Constants.Verbose {
			fmt.Println("start sorting")
		}
		newArr := make([]int, 2)
		if a[0] > b[0] {
			newArr[0] = b[0]
			newArr[1] = a[0]
		} else {
			newArr[0] = a[0]
			newArr[1] = b[0]
		}
		if Constants.Verbose {
			fmt.Println(newArr)
		}
		return newArr
	} else {
		if Constants.Verbose {
			fmt.Println("start splitting")
		}
		newArr := make([]int, lengthA+lengthB)
		left := MergeSort(a[:lengthA/2], a[lengthA/2:])
		right := MergeSort(b[:lengthB/2], b[lengthB/2:])
		if Constants.Verbose {
			fmt.Println("start merging")
			fmt.Println(left)
			fmt.Println(right)
		}
		ri := 0
		li := 0
		ai := 0
		for li < len(left) && ri < len(right) {
			if left[li] < right[ri] {
				newArr[ai] = left[li]
				li++
				ai++
			} else {
				newArr[ai] = right[ri]
				ri++
				ai++
			}
		}

		if li == len(left) {
			for ; ri < len(right); ri++ {
				newArr[ai] = right[ri]
				ai++
			}
		} else {
			for ; li < len(left); li++ {
				newArr[ai] = left[li]
				ai++
			}
		}

		return newArr
	}
}
