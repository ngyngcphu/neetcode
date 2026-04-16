func twoSum(nums []int, target int) []int {
	sort.Ints(nums)

	i, j := 0, len(nums) - 1
	for i < j {
		cur := nums[i] + nums[j]
		if cur == target {
			return []int{i, j}
		} else if cur < target {
			i++
		} else {
			j--
		}
	}
	return []int{}
}
