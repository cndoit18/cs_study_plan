package linkedlist

type MyLinkedList struct {
	root *ListNode
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		root: &ListNode{},
	}
}

func (this *MyLinkedList) Get(index int) int {
	i := 0
	current := this.root.Next
	for current != nil && i < index {
		current, i = current.Next, i+1
	}

	if current == nil {
		return -1
	}
	return current.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	node := &ListNode{Next: this.root.Next, Val: val}
	this.root.Next = node
}

func (this *MyLinkedList) AddAtTail(val int) {
	prev := this.root
	current := prev.Next
	for current != nil {
		prev, current = current, current.Next
	}
	prev.Next = &ListNode{Val: val}
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	prev := this.root
	current := prev.Next
	i := 0
	for current != nil && i < index {
		prev, current, i = current, current.Next, i+1
	}
	if i < index {
		return
	}
	node := &ListNode{Val: val, Next: current}
	prev.Next = node
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	prev := this.root
	current := prev.Next
	i := 0
	for current != nil && i < index {
		prev, current, i = current, current.Next, i+1
	}

	if current != nil {
		prev.Next = current.Next
	} else {
		prev.Next = nil
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
