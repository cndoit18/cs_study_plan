package stack

import (
	"bytes"
	"container/list"
)

func removeDuplicates(s string) string {
	stack := list.New()
	for _, v := range s {
		if stack.Len() != 0 && stack.Back().Value.(rune) == v {
			stack.Remove(stack.Back())
			continue
		}
		stack.PushBack(v)
	}
	buffer := bytes.Buffer{}
	for node := stack.Front(); node != nil; node = node.Next() {
		buffer.WriteRune(node.Value.(rune))
	}
	return buffer.String()
}
