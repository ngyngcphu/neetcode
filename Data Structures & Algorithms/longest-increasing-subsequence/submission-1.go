func lengthOfLIS(nums []int) int {
    arr := []int{}
    search := func(index int) int {
        left := 0
        right := len(arr) - 1
        for left < right {
            mid := (left + right) / 2
            if arr[mid] >= nums[index] {
                right = mid
            } else {
                left = mid + 1
            }
        }
        return right
    }
    for i := range nums {
        if i == 0 || nums[i] > arr[len(arr) - 1] {
            arr = append(arr, nums[i])
        } else {
            arr[search(i)] = nums[i]
        }
    }
    return len(arr)
}
