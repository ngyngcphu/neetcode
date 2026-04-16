func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	hm := make(map[int]int)
	for _, n := range nums {
		hm[n]++
	}

	res := [][]int{}
	for i := 0; i < n - 2; i++ {
		hm[nums[i]]--
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}
		for j := i + 1; j < n - 1; j++ {
			hm[nums[j]]--
			if j > i + 1 && nums[j] == nums[j - 1] {
				continue
			}
			target := -(nums[i] + nums[j])
			if hm[target] > 0 {
				res = append(res, []int{nums[i], nums[j], target})
			}
		}
		for j := i + 1; j < n - 1; j++ {
			hm[nums[j]]++
		}
	}
	return res
}
