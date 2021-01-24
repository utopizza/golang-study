package sorting

// ShellSort 希尔排序
func ShellSort(arr []int) {
	h := 1
	for 3*h < len(arr) {
		h = 3*h + 1
	}
	for h >= 1 {
		for i := range arr {
			for j := i; j >= h && arr[j] < arr[j-h]; j -= h {
				swap(arr, j, j-h)
			}
		}
		h = h / 3
	}
}
