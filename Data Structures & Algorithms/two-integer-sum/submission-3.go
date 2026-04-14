func twoSum(nums []int, target int) []int {
	count := make(map[int]int)

	for i, num := range nums {
		count[num] = i
	}
	for i, num := range nums {
		if j, found := count[target - num]; found && j != i {
			return []int{i, j}
		}
	}
	return []int{}
}
