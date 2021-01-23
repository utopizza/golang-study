package sorting

func bubbleSort(arr []int) {
	for i := len(arr) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				swap(arr, j, j+1)
			}
		}
	}
}
