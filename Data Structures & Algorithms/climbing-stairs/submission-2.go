func climbStairs(n int) int {
	prev, cur := 1, 1
	for i := 0; i <= n - 2; i++ {
		prev, cur = cur, prev + cur
	}
	return cur
}
