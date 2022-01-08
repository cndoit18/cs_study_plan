package stack

import (
	"container/list"
	"strconv"
)

func evalRPN(tokens []string) int {
	// 栈的先后顺序发生了变化
	operators := map[string]func(x, y int) int{
		"+": func(x, y int) int {
			return y + x
		},
		"-": func(x, y int) int {
			return y - x
		},
		"*": func(x, y int) int {
			return y * x
		},
		"/": func(x, y int) int {
			return y / x
		},
	}

	stack := list.New()
	for _, v := range tokens {
		if f, ok := operators[v]; ok {
			stack.PushBack(f(stack.Remove(stack.Back()).(int), stack.Remove(stack.Back()).(int)))
			continue
		}

		// 一定正确，忽略错误
		i, _ := strconv.Atoi(v)
		stack.PushBack(i)
	}
	return stack.Back().Value.(int)
}
