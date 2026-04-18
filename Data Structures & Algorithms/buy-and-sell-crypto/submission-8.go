func maxProfit(prices []int) int {
	if len(prices) == 0 {
        return 0
    }
	
	maxPrice := 0
	start := 0
	for end := 1; end < len(prices); end++ {
		if prices[end] < prices[start] {
			start = end
		}
		if maxPrice < prices[end] - prices[start] {
			maxPrice = prices[end] - prices[start]
		}
	}
	return maxPrice
}
