func minCostClimbingStairs(cost []int) int {
	memo := make(map[int]int, len(cost) + 1)
	var helper func(n int) int
	helper = func(n int) int {
		if n == 0 || n == 1 {
			return 0
		}

		if val, ok := memo[n]; ok {
			return val
		}

		memo[n] = min(helper(n - 1) + cost[n - 1], helper(n - 2) + cost[n - 2])
		return memo[n]
	}
	return helper(len(cost))
}
