func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make(map[int][]int)
	inDegree := make([]int, numCourses)
	for _, pre := range prerequisites {
		src, dest := pre[1], pre[0]
		graph[src] = append(graph[src], dest)
		inDegree[dest]++
	}
	queue := []int{}
	for i, d := range inDegree {
		if d == 0 {
			queue = append(queue, i)
		}
	}
	count := 0
	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]
		count++
		for _, e := range graph[course] {
			inDegree[e]--
			if inDegree[e] == 0 {
				queue = append(queue, e)
			}
		}
	}
	return count == numCourses
}
