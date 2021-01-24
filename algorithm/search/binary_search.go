package search

// BinarySearch 二分查找
func BinarySearch(arr []int, target int) int {
	for lo, hi := 0, len(arr)-1; lo <= hi; {
		mid := lo + (hi-lo)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return -1
}
