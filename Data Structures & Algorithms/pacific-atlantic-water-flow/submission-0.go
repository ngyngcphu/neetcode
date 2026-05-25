func pacificAtlantic(heights [][]int) [][]int {
	rows, cols := len(heights), len(heights[0])
	pacific := make(map[string]bool)
	atlantic := make(map[string]bool)
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	var dfs func(r, c int, visited map[string]bool)
	dfs = func(r, c int, visited map[string]bool) {
		key := fmt.Sprintf("%d,%d", r, c)
		if visited[key] {
			return
		}
		visited[key] = true
		for _, d := range directions {
			u, v := r + d[0], c + d[1]
			if u < 0 || u >= rows || v < 0 || v >= cols {
				continue
			}
			if heights[u][v] >= heights[r][c] {
				dfs(u, v, visited)
			}
		}
	}

	for r := range rows {
		dfs(r, 0, pacific)
		dfs(r, cols - 1, atlantic)
	}
	for c := range cols {
		dfs(0, c, pacific)
		dfs(rows - 1, c, atlantic)
	}

	result := [][]int{}
	for p := range pacific {
		if atlantic[p] {
			parts := strings.Split(p, ",")
			u, _ := strconv.Atoi(parts[0])
			v, _ := strconv.Atoi(parts[1])
			result = append(result, []int{u, v})
		}
	}
	return result
}
