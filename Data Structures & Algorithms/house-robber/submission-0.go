func rob(nums []int) int {
	memo := make(map[int]int, len(nums) + 1)
	var robhelper func(n int) int
	robhelper = func(n int) int {
		if n == 0 {
			return 0
		}
		if n == 1 {
			return nums[0]
		}

		if val, ok := memo[n]; ok {
			return val
		}
		memo[n] = max(robhelper(n - 2) + nums[n - 1], robhelper(n - 1))
		return memo[n]
	}
	return robhelper(len(nums))
}
