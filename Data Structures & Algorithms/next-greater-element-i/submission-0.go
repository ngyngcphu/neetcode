func nextGreaterElement(nums1 []int, nums2 []int) []int {
	nums1Index := make(map[int]int)
	for i, num := range nums1 {
		nums1Index[num] = i
	}
	res := make([]int, len(nums1))
	for i := range res {
		res[i] = -1
	}
	stack := []int{}

	for _, num := range nums2 {
		for len(stack) > 0 && num > nums1[stack[len(stack) - 1]] {
			idx := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
			res[idx] = num
		}
		if _, ok := nums1Index[num]; ok {
			stack = append(stack, nums1Index[num])
		}
	}
	return res
}
