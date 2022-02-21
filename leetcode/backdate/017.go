package backdate

var keyborad = map[string][]string{
	"1": {},
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	if len(digits) == 1 {
		return keyborad[digits]
	}

	cur := string(digits[0])
	l := letterCombinations(digits[1:])
	result := []string{}
	for _, v := range keyborad[cur] {
		for _, ll := range l {
			result = append(result, v+ll)
		}
	}
	return result
}
