func rob(nums []int) int {
	prev := 0
	cur := nums[0]
	for i := 2; i <= len(nums); i++ {
		prev, cur = cur, max(cur, prev + nums[i - 1])
	}
	return cur
}
