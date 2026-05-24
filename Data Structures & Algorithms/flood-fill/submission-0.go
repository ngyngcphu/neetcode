func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	rows, cols := len(image), len(image[0])
	orgColor := image[sr][sc]
	if orgColor == color {
		return image
	}
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= rows || c < 0 || c >= cols {
			return
		}
		if image[r][c] == orgColor {
			image[r][c] = color
			for _, d := range directions {
				dfs(r + d[0], c + d[1])
			}
		}
	}
	dfs(sr, sc)
	return image
}
