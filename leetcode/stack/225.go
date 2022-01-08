package stack

import "container/list"

type MyStack struct {
	linked *list.List
}

func Constructor() MyStack {
	return MyStack{
		linked: list.New(),
	}
}

func (this *MyStack) Push(x int) {
	this.linked.PushBack(x)
}

func (this *MyStack) Pop() int {
	backup := list.New()
	for node := this.linked.Front(); node != nil; node = node.Next() {
		if node.Next() == nil {
			defer this.linked.PushBackList(backup)
			return this.linked.Remove(node).(int)
		}
	}

	return -1
}

func (this *MyStack) Top() int {
	backup := list.New()
	for node := this.linked.Front(); node != nil; node = node.Next() {
		if node.Next() == nil {
			defer this.linked.PushBackList(backup)
			backup.PushBack(node.Value)
			return this.linked.Remove(node).(int)
		}
	}

	return -1
}

func (this *MyStack) Empty() bool {
	return this.linked.Len() == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
