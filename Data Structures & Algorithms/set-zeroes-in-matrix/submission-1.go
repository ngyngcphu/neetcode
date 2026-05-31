func setZeroes(matrix [][]int) {
	rows, cols := len(matrix), len(matrix[0])
	firstRowZero, firstColZero := false, false
	for i := range cols {
		if matrix[0][i] == 0 {
			firstRowZero = true
			break
		}
	}
	for i := range rows {
		if matrix[i][0] == 0 {
			firstColZero = true
			break
		}
	}
	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}
	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if firstRowZero {
		for i := range cols {
			matrix[0][i] = 0
		}
	}
	if firstColZero {
		for i := range rows {
			matrix[i][0] = 0
		}
	}
}
