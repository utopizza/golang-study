package leetcode

import "math"

func nextPermutation(nums []int) {
	// 1、find a[i-1] < a[i] from right to left
	i := len(nums) - 1
	for i > 0 {
		if nums[i-1] < nums[i] {
			break
		}
		i--
	}

	if i > 0 {
		// 2、find the min element which larger than a[i-1] from a[i] to right end
		minJ := i
		minV := math.MaxInt64
		for j := i; j < len(nums); j++ {
			if nums[j] > nums[i-1] && nums[j] <= minV { //等号也记录，是因为如果有相同元素，需要放到最右边的位置，保持右边区间逆序
				minJ = j
				minV = nums[j]
			}
		}

		// 3、swap a[i-1] with a[j]
		nums[i-1], nums[minJ] = nums[minJ], nums[i-1]
	}

	// 4、reverse from a[i] to right end
	for p, q := i, len(nums)-1; p < q; {
		nums[p], nums[q] = nums[q], nums[p]
		p++
		q--
	}
}
