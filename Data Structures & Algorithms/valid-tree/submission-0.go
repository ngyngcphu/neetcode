func validTree(n int, edges [][]int) bool {
	adjList := make(map[int][]int)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adjList[u] = append(adjList[u], v)
		adjList[v] = append(adjList[v], u)
	}
	visited := make([]bool, n)
	if hasCycle(adjList, visited, 0, -1) {
		return false
	}
	for _, v := range visited {
		if !v {
			return false
		}
	}
	return true
}

func hasCycle(adjList map[int][]int, visited []bool, node, parent int) bool {
	visited[node] = true
	for _, neighbor := range adjList[node] {
		if visited[neighbor] && neighbor != parent {
			return true
		}
		if !visited[neighbor] && hasCycle(adjList, visited, neighbor, node) {
			return true
		}
	}
	return false
}
