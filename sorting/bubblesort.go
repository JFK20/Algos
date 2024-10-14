package sorting

func BubbleSort(arr []int) {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
		}
	}
}
