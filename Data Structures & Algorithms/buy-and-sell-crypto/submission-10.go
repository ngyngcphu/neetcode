func maxProfit(prices []int) int {
	low, high := 0, 1
	result := 0
	for low < high && high < len(prices) {
		diff := prices[high] - prices[low]
		if diff <= 0 {
			low = high
		} else {
			result = max(result, diff)
		}
		high++
	}
	return result
}
