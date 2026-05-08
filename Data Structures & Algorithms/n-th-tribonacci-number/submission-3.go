func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	prev2, prev1, cur := 0, 1, 1
	for i := 3; i <= n; i++ {
		prev2, prev1, cur = prev1, cur, prev2 + prev1 + cur
	}
	return cur
}
