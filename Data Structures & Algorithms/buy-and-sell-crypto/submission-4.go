func maxProfit(prices []int) int {
	maxPrice := -100
	start := 0
	for end := 1; end < len(prices); end++ {
		for prices[end] - prices[start] < 0 && start < end {
			start++
		}
		if maxPrice < prices[end] - prices[start] {
			maxPrice = prices[end] - prices[start]
		}
	}
	if maxPrice == -100 {
		return 0
	} else {
		return maxPrice
	}
}
