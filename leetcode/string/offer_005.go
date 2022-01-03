package string

func replaceSpace(s string) string {
	space := 0
	for _, c := range s {
		if c == ' ' {
			space++
		}
	}
	// '%20' replace ' '
	result := append([]rune(s), make([]rune, space*2)...)
	right, left := len(s)-1, len(result)-1

	o := []rune("02%")
	for right != left {
		if result[right] != ' ' {
			result[left] = result[right]
			right, left = right-1, left-1
			continue
		}
		for _, c := range o {
			result[left] = c
			left--
		}
		right--
	}
	return string(result)
}
