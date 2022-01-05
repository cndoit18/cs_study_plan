package string

func repeatedSubstringPattern(s string) bool {
	if len(s) == 0 {
		return false
	}

	next := make([]int, len(s))
	for current, i := 1, 0; current < len(next); current++ {
		for i > 0 && s[i] != s[current] {
			i = next[i-1]
		}
		if s[i] == s[current] {
			i++
		}
		next[current] = i
	}

	return next[len(s)-1] > 0 && len(s)%(len(s)-next[len(s)-1]) == 0
}
