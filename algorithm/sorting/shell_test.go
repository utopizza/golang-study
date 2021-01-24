package sorting

import (
	"testing"
)

func TestShellSort(t *testing.T) {
	source1 := []int{6, 5, 4, 3, 2, 1}
	target1 := []int{1, 2, 3, 4, 5, 6}
	ShellSort(source1)
	if !equal(source1, target1) {
		t.Error(source1)
	}
}
