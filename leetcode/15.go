package leetcode

import "sort"

func threeSum(nums []int) [][]int {
	var results [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i == 0 || nums[i] != nums[i-1] {
			target := -nums[i]
			for lo, hi := i+1, len(nums)-1; lo < hi; {
				if nums[lo]+nums[hi] == target {
					results = append(results, []int{nums[i], nums[lo], nums[hi]})
					for ; lo < hi && nums[lo+1] == nums[lo]; lo++ {
					}
					for ; lo < hi && nums[hi-1] == nums[hi]; hi-- {
					}
					lo++
					hi--
				} else if nums[lo]+nums[hi] > target {
					hi--
				} else {
					lo++
				}
			}
		}
	}
	return results
}
