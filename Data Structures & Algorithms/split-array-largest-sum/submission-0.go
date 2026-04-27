func splitArray(nums []int, k int) int {
	left, right := 0, 0
	for _, num := range nums {
		if left < num {
			left = num
		}
		right += num
	}

	canSplit := func(sum int) bool {
		subset := 1
		curSum := 0
		for _, num := range nums {
			if curSum + num <= sum {
				curSum += num
			} else {
				curSum = num
				subset += 1
			}
		}
		return subset <= k
	}

	for left < right {
		mid := (left + right) / 2
		if canSplit(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
