func twoSum(nums []int, target int) []int {
	newA := make([][2]int, len(nums))
	for i, num := range nums {
		newA[i] = [2]int{num, i}
	}
	sort.Slice(newA, func(i, j int) bool {
		return newA[i][0] < newA[j][0]
	})

	i, j := 0, len(nums) - 1
	for i < j {
		cur := newA[i][0] + newA[j][0]
		if cur == target {
			if newA[i][1] < newA[j][1] {
				return []int{newA[i][1], newA[j][1]}
			} else {
				return []int{newA[j][1], newA[i][1]}
			}
		} else if cur < target {
			i++
		} else {
			j--
		}
	}
	return []int{}
}
