func searchMatrix(matrix [][]int, target int) bool {
	row := len(matrix) - 1
	col := 0
	for row >= 0 && col < len(matrix[0]) {
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] < target {
			col++
		} else {
			row--
		}
	}
	return false
}
