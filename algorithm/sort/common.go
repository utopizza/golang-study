package sorting

import (
	"math/rand"
	"time"
)

func swap(source *[]int, i int, j int) {
	if i > len(*source) || j > len(*source) {
		return
	}
	(*source)[i], (*source)[j] = (*source)[j], (*source)[i]
}

func equal(source []int, target []int) bool {
	if len(source) != len(target) {
		return false
	}
	for i := range source {
		if source[i] != target[i] {
			return false
		}
	}
	return true
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func shuffle(source *[]int) {
	if len(*source) == 0 {
		return
	}
	rand.Seed(time.Now().Unix())
	for i := range *source {
		newIndex := rand.Intn(len(*source))
		swap(source, i, newIndex)
	}
}
