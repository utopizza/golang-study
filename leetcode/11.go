package leetcode

func maxArea(height []int) int {
	var maxArea int
	for i, j := 0, len(height)-1; i < j; {
		maxArea = max(maxArea, (j-i)*min(height[i], height[j]))
		if height[i] <= height[j] {
			i++
		} else {
			j--
		}
	}
	return maxArea
}
