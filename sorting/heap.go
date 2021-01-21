package sorting

func heapSort(source *[]int) {
	if len(*source) == 0 {
		return
	}
	N := len(*source)

	heap := make([]int, N+1)
	for i := range *source {
		heap[i+1] = (*source)[i]
	}

	for k := N / 2; k >= 1; k-- {
		sink(&heap, k, N)
	}

	for N > 1 {
		swap(&heap, 1, N)
		N--
		sink(&heap, 1, N)
	}

	for i := range *source {
		(*source)[i] = heap[i+1]
	}
}

func sink(heap *[]int, k, N int) {
	for k*2 <= N {
		j := k * 2
		if j < N && (*heap)[j] < (*heap)[j+1] {
			j++
		}
		if (*heap)[k] >= (*heap)[j] {
			break
		}
		swap(heap, k, j)
		k = j
	}
}
