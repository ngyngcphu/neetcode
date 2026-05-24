/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	oldToNew := make(map[*Node]*Node)
	return dfs(node, oldToNew)
}

func dfs(node *Node, oldToNew map[*Node]*Node) *Node {
	if node == nil {
		return nil
	}
	if _, found := oldToNew[node]; found {
		return oldToNew[node]
	}
	copy := &Node{Val: node.Val}
	oldToNew[node] = copy
	for _, neighbor := range node.Neighbors {
		copy.Neighbors = append(copy.Neighbors, dfs(neighbor, oldToNew))
	}
	return copy
}
