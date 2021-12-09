package linkedlist

func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	for node := dummy; node != nil && node.Next != nil; {
		for node.Next != nil && node.Next.Val == val {
			node.Next = node.Next.Next
		}
		node = node.Next
	}
	return dummy.Next
}
