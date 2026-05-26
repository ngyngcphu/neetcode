func combinationSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	path := []int{}
	var dfs func(int, int)
	dfs = func(start, target int) {
		if target == 0 {
			result = append(result, append([]int{}, path...))
			return
		}
		for i := start; i < len(nums); i++ {
			if nums[i] > target {
				return
			}
			path = append(path, nums[i])
			dfs(i, target - nums[i])
			path = path[:len(path) - 1]
		}
	}
	dfs(0, target)
	return result
}
