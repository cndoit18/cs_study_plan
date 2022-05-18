package greedy

import "strconv"

func monotoneIncreasingDigits(n int) int {
	s := []byte(strconv.Itoa(n))
	for i := len(s) - 1; i > 0; i-- {
		if s[i] < s[i-1] {
			for j := i; j < len(s); j++ {
				s[j] = '9'
			}
			s[i-1]--
		}
	}
	r, _ := strconv.Atoi(string(s))
	return r
}
