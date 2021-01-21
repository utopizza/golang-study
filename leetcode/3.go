package leetcode

func lengthOfLongestSubstring(s string) int {
	charArr := []rune(s)
	charMap := make(map[rune]int)
	var left, maxLen int
	for i := 0; i < len(charArr); i++ {
		c := charArr[i]
		// 如果发生重复，则向右移动left指针直到没有重复为止，相当于dequeue当前窗口
		if lastSeenIndex, ok := charMap[c]; ok {
			left = max(left, lastSeenIndex+1)
		}
		// 更新当前窗口相关值
		maxLen = max(maxLen, i-left+1)
		charMap[c] = i
	}
	return maxLen
}
