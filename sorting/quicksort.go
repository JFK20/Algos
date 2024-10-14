package sorting

import "sync"

func QuickSort(arr []int, low int, high int) {
	if low < high {
		pivot := partition(arr, low, high)

		QuickSort(arr, low, pivot-1)
		QuickSort(arr, pivot+1, high)
	}
}

func partition(arr []int, low int, high int) int {
	pivot := arr[high]

	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func QuickSortAsync(arr []int, low int, high int) {
	if low < high {
		pivot := partition(arr, low, high)

		if len(arr) > 128 {
			var wg sync.WaitGroup
			wg.Add(2)
			go func() {
				QuickSortAsync(arr, low, pivot-1)
				wg.Done()
			}()
			go func() {
				QuickSortAsync(arr, pivot+1, high)
				wg.Done()
			}()

			wg.Wait()
		} else {
			QuickSort(arr, low, pivot-1)
			QuickSort(arr, pivot+1, high)
		}
	}
}
