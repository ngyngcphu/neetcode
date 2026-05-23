/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
	dia := 0
	dfs(root, &dia)
	return dia
}

func dfs(node *TreeNode, dia *int) int {
	if node == nil {
		return 0
	}
	left := dfs(node.Left, dia)
	right := dfs(node.Right, dia)
	*dia = max(*dia, left + right)
	return 1 + max(left, right)
}
