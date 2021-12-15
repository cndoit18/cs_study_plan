package linkedlist

// 给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{
		Next: head,
	}

	current := dummy
	for current.Next != nil && current.Next.Next != nil {
		var tmp *ListNode
		current.Next, tmp = current.Next.Next, current.Next
		current.Next.Next, tmp.Next = tmp, current.Next.Next
		current = tmp
	}
	return dummy.Next
}
