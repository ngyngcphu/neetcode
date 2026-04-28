func minCostClimbingStairs(cost []int) int {
	prev, cur := 0, 0
	for i := 2; i <= len(cost); i++ {
		prev, cur = cur, min(prev + cost[i - 2], cur + cost[i - 1])
	}
	return cur
}
