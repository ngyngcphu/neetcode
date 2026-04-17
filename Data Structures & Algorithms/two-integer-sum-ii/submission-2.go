func twoSum(numbers []int, target int) []int {
	for i, n := range numbers {
		l, r := i + 1, len(numbers) - 1
		diff := target - n
		for l <= r {
			mid := (l + r) / 2
			if numbers[mid] == diff {
				return []int{i + 1, mid + 1}
			} else if numbers[mid] < diff {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return []int{}
}