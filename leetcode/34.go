package leetcode

func searchRange(nums []int, target int) []int {
	return []int{findLeft(nums, target), findRight(nums, target)}
}

func findLeft(nums []int, target int) int {
	if nums != nil && len(nums) > 0 {
		var lo, hi = 0, len(nums) - 1
		for lo < hi {
			mid := (lo + hi) / 2
			if nums[mid] >= target {
				hi = mid
			} else {
				lo = mid + 1
			}
		}
		if nums[hi] == target {
			return hi
		}
	}
	return -1
}

func findRight(nums []int, target int) int {
	if nums != nil && len(nums) > 0 {
		var lo, hi = 0, len(nums) - 1
		for lo < hi {
			mid := (lo + hi + 1) / 2
			if nums[mid] <= target {
				lo = mid
			} else {
				hi = mid - 1
			}
		}
		if nums[lo] == target {
			return lo
		}
	}
	return -1
}
