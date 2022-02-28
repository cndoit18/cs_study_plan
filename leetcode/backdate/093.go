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

func split(prev []string, left string, result *[]string) {
	if len(prev) == 3 && len(left) > 0 {
		c, _ := strconv.Atoi(left)
		if !(c <= 255) || (len(left) > 1 && left[0] == '0') {
			return
		}
		prev = append(prev, left)
		*result = append(*result, strings.Join(prev, "."))
		return
	}
	c := 0
	for i := 0; i < 4 && len(left) > i; i++ {
		tmp := 10 * c
		tmp += int(left[i] - byte('0'))
		if tmp > 255 || (i > 0 && left[0] == '0') {
			break
		}
		c = tmp
		sub := ""
		if i+1 <= len(left) {
			sub = left[i+1:]
		}
		s := append([]string{}, prev...)
		s = append(s, left[:i+1])
		split(s, sub, result)
	}
}
