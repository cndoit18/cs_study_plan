package string

func reverseWords(s string) string {
	inplace := []rune(s)
	fast, slow := 0, 0
	for fast < len(inplace) {
		if inplace[fast] == ' ' && (fast == 0 || inplace[fast-1] == ' ') {
			fast++
			continue
		}
		inplace[slow] = inplace[fast]
		fast, slow = fast+1, slow+1
	}

	if inplace[slow-1] == ' ' {
		slow--
	}
	inplace = inplace[:slow]

	for i, j := 0, len(inplace)-1; i < j; i, j = i+1, j-1 {
		inplace[i], inplace[j] = inplace[j], inplace[i]
	}

	for fast, slow = 0, 0; fast <= len(inplace); {
		if fast != len(inplace) && inplace[fast] != ' ' {
			fast++
			continue
		}
		for i, j := slow, fast-1; i < j; i, j = i+1, j-1 {
			inplace[i], inplace[j] = inplace[j], inplace[i]
		}
		slow, fast = fast+1, fast+1
	}
	return string(inplace)
}
