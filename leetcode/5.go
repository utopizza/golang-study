package leetcode

func longestPalindrome(s string) string {
	arr := []rune(s)
	var currentStr, longestStr string
	for i := 0; i < len(s); i++ {
		// even
		currentStr = expend(arr, i, i+1)
		if len(currentStr) > len(longestStr) {
			longestStr = currentStr
		}
		// odd
		currentStr = expend(arr, i-1, i+1)
		if len(currentStr) > len(longestStr) {
			longestStr = currentStr
		}
	}
	return longestStr
}

func expend(arr []rune, left, right int) string {
	for left >= 0 && right < len(arr) {
		if arr[left] == arr[right] {
			left--
			right++
		} else {
			break
		}
	}
	return string(arr[left+1 : right])
}
