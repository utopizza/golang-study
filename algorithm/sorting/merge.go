package sorting

// MergeSort 归并排序
func MergeSort(arr []int) {
	mergeSortRecur(arr, 0, len(arr)-1)
}

func mergeSortRecur(arr []int, lo, hi int) {
	if lo >= hi {
		return
	}
	mid := lo + (hi-lo)/2
	mergeSortRecur(arr, lo, mid)
	mergeSortRecur(arr, mid+1, hi)
	merge(arr, lo, mid, hi)
}

func merge(arr []int, lo, mid, hi int) {
	// 拷贝到一个临时数组
	temp := make([]int, len(arr))
	for k := lo; k <= hi; k++ {
		temp[k] = arr[k]
	}

	// 比较并移动两个头指针
	i, j := lo, mid+1
	for k := lo; k <= hi; k++ {
		if i > mid { // 左半部分已经全部出列
			arr[k] = temp[j]
			j++
		} else if j > hi { // 右半部分已全部出列
			arr[k] = temp[i]
			i++
		} else if temp[i] <= temp[j] { // 取较小的队头出列
			arr[k] = temp[i]
			i++
		} else {
			arr[k] = temp[j]
			j++
		}
	}
}
