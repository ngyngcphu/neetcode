/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := [][]int{}
	queue := []*TreeNode{root}
	leftToRight := true
	for len(queue) > 0 {
		levelSize := len(queue)
		curLevel := []int{}
		for range levelSize {
			node := queue[0]
			queue = queue[1:]
			if leftToRight {
				curLevel = append(curLevel, node.Val)
			} else {
				curLevel = append([]int{node.Val}, curLevel...)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		leftToRight = !leftToRight
		result = append(result, curLevel)
	}
	return result
}
