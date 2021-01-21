package sorting

func quickSort(source *[]int) {
	shuffle(source)
	quickSortRecur(source, 0, len(*source)-1)
}

func quickSortRecur(source *[]int, lo, hi int) {
	if lo >= hi {
		return
	}
	mid := quickSortPartition(source, lo, hi)
	quickSortRecur(source, lo, mid-1)
	quickSortRecur(source, mid+1, hi)
}

func quickSortPartition(source *[]int, lo, hi int) int {
	midIndex := lo
	midValue := (*source)[midIndex]
	i, j := lo, hi
	for {
		for ; (*source)[i] <= midValue; i++ {
			if i == hi {
				break
			}
		}
		for ; (*source)[j] >= midValue; j-- {
			if j == lo {
				break
			}
		}
		if i >= j {
			break
		}
		swap(source, i, j)
	}
	swap(source, midIndex, j)
	return j
}
