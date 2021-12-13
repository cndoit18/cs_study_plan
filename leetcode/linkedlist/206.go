package linkedlist

// 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
func reverseList(head *ListNode) *ListNode {
	var prev, current *ListNode = nil, head
	for current != nil {
		current.Next, prev, current = prev, current, current.Next
	}
	return prev
}
