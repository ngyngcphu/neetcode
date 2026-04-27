func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	robRange := func(start, end int) int {
		memo := make(map[int]int, n)

		var helper func(n int) int
		helper = func(n int) int {
			if n == start {
				return 0
			}
			if n == start + 1 {
				return nums[start]
			}

			if val, ok := memo[n]; ok {
				return val
			}
			memo[n] = max(helper(n - 1), helper(n - 2) + nums[n - 1])
			return memo[n]
		}
		return helper(end)
	}

	return max(robRange(0, n - 1), robRange(1, n))
}
