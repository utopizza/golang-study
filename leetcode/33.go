package leetcode

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	p := findPivot(nums)
	if ret := binarySearch(nums, 0, p-1, target); ret != -1 {
		return ret
	}
	if ret := binarySearch(nums, p, len(nums)-1, target); ret != -1 {
		return ret
	}
	return -1
}

func findPivot(nums []int) int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := (lo + hi) / 2
		if nums[mid] < nums[hi] { //说明mid到hi是升序，pivot不在此区间
			hi = mid
		} else { //说明mid到hi是乱序，pivot必在此区间内
			lo = mid + 1
		}
	}
	return hi
}

func binarySearch(nums []int, lo, hi, target int) int {
	for lo <= hi {
		mid := (lo + hi) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return -1
}
