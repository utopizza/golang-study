package sorting

func selectionSort(source *[]int) {
	for i := range *source {
		min := i
		for j := i + 1; j < len(*source); j++ {
			if (*source)[j] < (*source)[min] {
				min = j
			}
		}
		swap(source, i, min)
	}
}
