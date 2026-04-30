func searchMatrix(matrix [][]int, target int) bool {
	top := 0
	bot := len(matrix) - 1
	for top <= bot {
		mid := (top + bot) / 2
		if matrix[mid][0] > target {
			bot = mid - 1
		} else if matrix[mid][len(matrix[0]) - 1] < target {
			top = mid + 1
		} else {
			break
		}
	}
	row := (top + bot) / 2
	
	left := 0
	right := len(matrix[0]) - 1
	for left <= right {
		col := (left + right) / 2
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] < target {
			left = col + 1
		} else {
			right = col - 1
		}
	}
	return false
}
