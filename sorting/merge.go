package sorting

func mergeSort(source *[]int) {
	aux := make([]int, len(*source))
	for step := 1; step <= len(*source); step *= 2 {
		for lo := 0; lo < len(*source)-step; lo += step * 2 {
			merge(source, &aux, lo, lo+step, min(lo+step*2-1, len(*source)-1))
		}
	}
}

func merge(source, aux *[]int, lo, mid, hi int) {
	if lo >= mid || mid > hi || hi >= len(*source) {
		return
	}
	i, j := lo, mid
	for k := lo; k <= hi; k++ {
		if i > mid {
			(*aux)[k] = (*source)[j]
			j++
		} else if j > hi {
			(*aux)[k] = (*source)[i]
			i++
		} else if (*source)[i] <= (*source)[j] {
			(*aux)[k] = (*source)[i]
			i++
		} else {
			(*aux)[k] = (*source)[j]
			j++
		}
	}
	for k := lo; k <= hi; k++ {
		(*source)[k] = (*aux)[k]
	}
}
