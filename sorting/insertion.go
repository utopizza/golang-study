package sorting

func insertionSort(source *[]int) {
	for i := range *source {
		for j := i; j > 0; j-- {
			if (*source)[j] < (*source)[j-1] {
				swap(source, j, j-1)
			}
		}
	}
}
