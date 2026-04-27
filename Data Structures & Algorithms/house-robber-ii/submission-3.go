func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	robRange := func(start, end int) int {
		dp := make([]int, end + 1)
		dp[start], dp[start + 1] = 0, nums[start]
		for i := start + 2; i <= end; i++ {
			dp[i] = max(dp[i - 1], dp[i - 2] + nums[i - 1])
		}
		return dp[end]
	}
	return max(robRange(0, n - 1), robRange(1, n))
}
