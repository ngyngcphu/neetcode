func tribonacci(n int) int {
	memo := make(map[int]int)
	var helper func(n int) int
	helper = func(n int) int {
		if n == 0 {
			return 0
		}
		if n == 1 || n == 2 {
			return 1
		}
		if val, ok := memo[n]; ok {
			return val
		}
		memo[n] = helper(n - 3) + helper(n - 2) + helper(n - 1)
		return memo[n]
	}
	return helper(n)
}
