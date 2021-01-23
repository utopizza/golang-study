package sorting

func selectionSort(arr []int) {
	for i := range arr {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		swap(arr, i, min)
	}
}
