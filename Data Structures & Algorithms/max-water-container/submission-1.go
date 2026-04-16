func maxArea(heights []int) int {
	left, right := 0, len(heights) - 1
	res := 0
	for left <= right {
		area := (right - left) * min(heights[left], heights[right])
		if res < area {
			res = area
		}
		if heights[left] < heights[right] {
			left++
		} else {
			right--
		}
	}
	return res
}
