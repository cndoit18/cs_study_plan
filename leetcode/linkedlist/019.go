package linkedlist

// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fast := 0
	dummy := &ListNode{
		Next: head,
	}
	current := dummy
	slowNode := dummy
	for current != nil {
		if fast <= n {
			current, fast = current.Next, fast+1
			continue
		}
		slowNode, current = slowNode.Next, current.Next
	}
	if slowNode.Next != nil {
		slowNode.Next = slowNode.Next.Next
	}

	return dummy.Next
}
