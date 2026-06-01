func maxProfit(prices []int) int {
	minPrice := prices[0]
	maxPrice := 0
	for i := 1; i < len(prices); i++ {
		maxPrice = max(maxPrice, prices[i] - minPrice)
		minPrice = min(minPrice, prices[i])
	}
	return maxPrice
}
