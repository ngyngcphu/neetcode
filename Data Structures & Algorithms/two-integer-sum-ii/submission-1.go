func twoSum(numbers []int, target int) []int {
	store := make(map[int]int)
	for i, n := range numbers {
		diff := target - n
		if store[diff] > 0 {
			return []int{store[diff], i + 1}
		}
		store[n] = i + 1
	}
	return []int{}
}
