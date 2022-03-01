package backdate

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	result := []string{}
	split([]string{}, s, &result)
	return result
}

func convert(num string) (int, bool) {
	if len(num) > 1 && num[0] == '0' {
		return 0, false
	}

	r, err := strconv.Atoi(num)
	return r, err == nil
}

func split(prev []string, left string, result *[]string) {
	if len(prev) == 3 && len(left) > 0 {
		if r, ok := convert(left); ok && r <= 255 {
			*result = append(*result, strings.Join(append(prev, left), "."))
		}
		return
	}
	for i := 1; len(left) >= i && i <= 4; i++ {
		r, ok := convert(left[:i])
		if !ok || !(r <= 255) {
			break
		}
		split(append(append([]string{}, prev...), left[:i]), left[i:], result)
	}
}
