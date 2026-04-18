func maxProfit(prices []int) int {
	l, r := 0, 1
	maxP := 0
	for r < len(prices) {
		diff := prices[r] - prices[l]
		if diff > 0 { 
			if maxP < diff {
				maxP = diff
			}
		} else {
			l = r
		}
		r++
	}
	return maxP
}
