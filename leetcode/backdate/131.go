package backdate

func partition(s string) [][]string {
	result := [][]string{}
	if len(s) == 0 {
		return result
	}

	if isPartition(s) {
		result = append(result, []string{s})
	}

	for i := 1; i < len(s); i++ {
		prefix := s[:i]
		if !isPartition(prefix) {
			continue
		}
		l := partition(s[i:])
		for _, v := range l {
			layer := append([]string{prefix}, v...)
			result = append(result, layer)
		}
	}

	return result
}

func isPartition(s string) bool {
	for r, l := len(s)-1, 0; r >= l; r, l = r-1, l+1 {
		if s[r] != s[l] {
			return false
		}
	}
	return true
}
