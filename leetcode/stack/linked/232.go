package linked

import (
	"container/list"
)

type MyQueue struct {
	inputStack  *list.List
	outputStack *list.List
}

func Constructor() MyQueue {
	return MyQueue{
		inputStack:  list.New(),
		outputStack: list.New(),
	}
}

func (this *MyQueue) Push(x int) {
	this.inputStack.PushFront(x)
}

func (this *MyQueue) Pop() int {
	if this.outputStack.Len() == 0 {
		for current := this.inputStack.Front(); current != nil; {
			next := current.Next()
			this.outputStack.PushFront(this.inputStack.Remove(current))
			current = next
		}
	}

	return this.outputStack.Remove(this.outputStack.Front()).(int)
}

func (this *MyQueue) Peek() int {
	if this.outputStack.Len() == 0 {
		for current := this.inputStack.Front(); current != nil; {
			next := current.Next()
			this.outputStack.PushFront(this.inputStack.Remove(current))
			current = next
		}
	}

	return this.outputStack.Front().Value.(int)
}

func (this *MyQueue) Empty() bool {
	return this.inputStack.Len() == 0 && this.outputStack.Len() == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
