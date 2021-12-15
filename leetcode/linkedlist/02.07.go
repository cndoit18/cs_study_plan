package linkedlist

// 给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表没有交点，返回 null 。
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	headALen, headBLen := 0, 0
	for current := headA; current != nil; current = current.Next {
		headALen++
	}
	for current := headB; current != nil; current = current.Next {
		headBLen++
	}

	for headALen > headBLen {
		headA = headA.Next
		headALen--
	}
	for headBLen > headALen {
		headB = headB.Next
		headBLen--
	}

	for headA != nil && headB != nil && headA != headB {
		headA = headA.Next
		headB = headB.Next
	}

	return headA
}
