func spiralOrder(matrix [][]int) []int {
	result := []int{}
	for len(matrix) > 0 {
		result = append(result, matrix[0]...)
		matrix = matrix[1:]
		if len(matrix) > 0 && len(matrix[0]) > 0 {
			size := len(matrix[0])
			for i := range matrix {
				result = append(result, matrix[i][size - 1])
				matrix[i] = matrix[i][:size - 1]
			}
		}
		if len(matrix) > 0 {
			lastRow := matrix[len(matrix) - 1]
			matrix = matrix[:len(matrix) - 1]
			for i := len(lastRow) - 1; i >= 0; i-- {
				result = append(result, lastRow[i])
			}
		}
		if len(matrix) > 0 && len(matrix[0]) > 0 {
			for i := len(matrix) - 1; i >= 0; i-- {
				result = append(result, matrix[i][0])
				matrix[i] = matrix[i][1:]
			}
		}
	}
	return result
}
