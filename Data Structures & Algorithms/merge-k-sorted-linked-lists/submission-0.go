/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {
	merge := func(l1, l2 *ListNode) *ListNode {
		dummy := &ListNode{}
		cur := dummy
		for l1 != nil && l2 != nil {
			if l1.Val <= l2.Val {
				cur.Next = l1
				l1 = l1.Next
			} else {
				cur.Next = l2
				l2 = l2.Next
			}
			cur = cur.Next
		}
		if l1 == nil {
			cur.Next = l2
		} else if l2 == nil {
			cur.Next = l1
		}
		return dummy.Next
	}

	var divide func(left, right int) *ListNode
	divide = func(left, right int) *ListNode {
		if left > right {
			return nil
		}
		if left == right {
			return lists[left]
		}
		mid := (left + right) / 2
		l1 := divide(left, mid)
		l2 := divide(mid + 1, right)
		return merge(l1, l2)
	}

	return divide(0, len(lists) - 1)
}

