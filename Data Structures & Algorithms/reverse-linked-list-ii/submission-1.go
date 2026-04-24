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
	cur := leftNode

	var prev *ListNode = nil
	for i < right {
		nextNode := cur.Next
		cur.Next = prev
		prev = cur
		cur = nextNode
		i++
	}
	prevLeft.Next = prev
	leftNode.Next = cur
	return dummy.Next
}
