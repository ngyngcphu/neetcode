/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	var prev *ListNode = nil
	for slow != nil {
		next := slow.Next
		slow.Next = prev
		prev = slow
		slow = next
	}
	cur := head
	for prev.Next != nil {
		temp := cur.Next
		cur.Next = prev
		cur = temp
		temp = prev.Next
		prev.Next = cur
		prev = temp
	}
}
