func twoSum(nums []int, target int) []int {
	prevMap := make(map[int]int)
	
	for i, num := range nums {
		if j, found := prevMap[target - num]; found {
			return []int{j, i}
		}
		prevMap[num] = i
	}
	return []int{}
}
