func subsets(nums []int) [][]int {
	result := [][]int{}
	var dfs func([]int, int)
	dfs = func(path []int, index int) {
		if index == len(nums) {
			result = append(result, append([]int{}, path...))
			return
		}
		path = append(path, nums[index])
		dfs(path, index + 1)
		path = path[:len(path) - 1]
		dfs(path, index + 1)
	}
	dfs([]int{}, 0)
	return result
}
