package leetcode

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sum := 0
	diff := math.MaxInt8
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		t := target - nums[i]
		for lo, hi := i+1, len(nums)-1; lo < hi; {
			if abs(t-nums[lo]-nums[hi]) < diff {
				diff = abs(t - nums[lo] - nums[hi])
				sum = nums[i] + nums[lo] + nums[hi]
			}
			if nums[lo]+nums[hi] == t {
				return target
			} else if nums[lo]+nums[hi] > t {
				hi--
			} else {
				lo++
			}
		}
	}
	return sum
}
