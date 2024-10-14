package main

import (
	"Algos/sorting"
	"fmt"
)

func main() {
	test := []int{17, 5, 420, 69, 187, 45, 800, 57, 7923, 23, 3, 99, 439, 123, 136, 45676}
	test = sorting.MergeSort(test[:len(test)/2], test[len(test)/2:])
	fmt.Println(test)
}
