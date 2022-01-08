package stack

import "container/list"

func isValid(s string) bool {
	left := list.New()
	match := map[rune]rune{
		'{': '}',
		'(': ')',
		'[': ']',
	}
	for _, v := range s {
		if _, ok := match[v]; ok {
			left.PushBack(v)
			continue
		}

		if left.Len() == 0 {
			return false
		}

		key := left.Remove(left.Back()).(rune)
		if match[key] != v {
			return false
		}
	}
	return left.Len() == 0
}
