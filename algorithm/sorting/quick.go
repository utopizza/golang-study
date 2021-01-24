package sorting

// QuickSort 快速排序
func QuickSort(arr []int) {
	shuffle(arr)
	quickSortRecur(arr, 0, len(arr)-1)
}

func quickSortRecur(arr []int, lo, hi int) {
	if lo >= hi {
		return
	}
	mid := quickSortPartition(arr, lo, hi)
	quickSortRecur(arr, lo, mid-1)
	quickSortRecur(arr, mid+1, hi)
}

func quickSortPartition(arr []int, lo, hi int) int {
	midIndex := lo
	midValue := arr[midIndex]
	i, j := lo+1, hi
	for {
		// 从左往右找，第一个比切分元素大的元素
		for ; i < hi && arr[i] <= midValue; i++ {
		}
		// 从右往左找，第一个比切分元素小的元素
		for ; j > lo && arr[j] >= midValue; j-- {
		}
		if i >= j {
			break
		}
		// 交换，保证比切分元素小的都在左边，大的都在右边
		swap(arr, i, j)
	}
	// 放置切分元素在合适的位置
	swap(arr, midIndex, j)
	return j
}
