package linkedlist

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	var meet *ListNode = nil
	for fast != nil && fast.Next != nil && slow != nil {
		fast, slow = fast.Next.Next, slow.Next
		if fast == slow {
			meet = fast
			break
		}
	}
	slow, fast = meet, head
	for slow != nil && fast != nil {
		if fast == slow {
			return fast
		}
		slow, fast = slow.Next, fast.Next
	}
	return nil
}
