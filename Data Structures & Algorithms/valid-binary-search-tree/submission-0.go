/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
	return dfs(root, math.MinInt32, math.MaxInt32)
}

func dfs(node *TreeNode, minVal, maxVal int) bool {
	if node == nil {
		return true
	}
	if node.Val <= minVal || node.Val >= maxVal {
		return false
	}
	return dfs(node.Left, minVal, node.Val) && dfs(node.Right, node.Val, maxVal)
}
