func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1
	minVal := 1000
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] >= nums[left] {
			minVal = min(minVal, nums[left])
			left = mid + 1
		} else {
			minVal = min(nums[mid], minVal)
			right = mid - 1
		}
	}
	return minVal
}
