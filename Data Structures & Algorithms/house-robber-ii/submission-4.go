func rob(nums []int) int {
    n := len(nums)
	if n == 1 {
		return nums[0]
	}

	robRange := func(start, end int) int {
		prev, cur := 0, nums[start]
		for i := start + 2; i <= end; i++ {
			prev, cur = cur, max(cur, prev + nums[i - 1])
		}
		return cur
	}
	return max(robRange(0, n - 1), robRange(1, n))
}
