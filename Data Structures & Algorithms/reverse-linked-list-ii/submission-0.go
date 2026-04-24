/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}
	prevLeft := dummy
	i := 0
	for i < left - 1 {
		prevLeft = prevLeft.Next
		i++
	}
	leftNode := prevLeft.Next
	i++
	rightNode := leftNode
	for i < right {
		rightNode = rightNode.Next
		i++
	}
	nextRight := rightNode.Next
	rightNode.Next = nil

	cur := leftNode
	var prev *ListNode = nil
	for cur != nil {
		temp := cur.Next
		cur.Next = prev
		prev = cur
		cur = temp
	}
	prevLeft.Next = prev
	leftNode.Next = nextRight
	return dummy.Next
}
