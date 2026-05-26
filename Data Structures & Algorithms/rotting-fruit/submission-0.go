func orangesRotting(grid [][]int) int {
	row, col := len(grid), len(grid[0])
	type Point [2]int
	var queue []Point
	lenFresh := 0
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	min := 0

	for r := range row {
		for c := range col {
			if grid[r][c] == 2 {
				queue = append(queue, Point{r, c})
			} else if grid[r][c] == 1 {
				lenFresh++
			}
		}
	}
	for len(queue) > 0 && lenFresh > 0 {
		levelSize := len(queue)
		for i := range levelSize {
			r, c := queue[i][0], queue[i][1]
			for _, d := range directions {
				nr, nc := r + d[0], c + d[1]
				if nr < 0 || nr >= row || nc < 0 || nc >= col || grid[nr][nc] != 1 {
					continue
				}
				grid[nr][nc] = 2
				lenFresh--
				queue = append(queue, Point{nr, nc})
			}
		}
		queue = queue[levelSize:]
		min++
	}
	if lenFresh == 0 {
		return min
	}
	return -1
}
