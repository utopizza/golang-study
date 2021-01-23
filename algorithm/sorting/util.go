package sorting

import (
	"math/rand"
	"time"
)

func swap(arr []int, i int, j int) {
	if i > len(arr) || j > len(arr) {
		return
	}
	arr[i], arr[j] = arr[j], arr[i]
}

func equal(arr []int, target []int) bool {
	if len(arr) != len(target) {
		return false
	}
	for i := range arr {
		if arr[i] != target[i] {
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

func shuffle(arr []int) {
	if len(arr) == 0 {
		return
	}
	rand.Seed(time.Now().Unix())
	for i := range arr {
		newIndex := rand.Intn(len(arr))
		swap(arr, i, newIndex)
	}
}
