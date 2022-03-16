package backdate

func solveNQueens(n int) [][]string {
	var f func(m []string, result *[][]string)
	f = func(m []string, result *[][]string) {
		if len(m) == n && len(m[0]) == n {
			*result = append(*result, m)
		}
		line := make([]byte, 0, n)
		for i := 0; i < n; i++ {
			line = append(line, '.')
		}
	outer:
		for i := 0; i < n; i++ {
			deep := append([]byte{}, line...)
			deep[i] = 'Q'
			for j := len(m) - 1; j > -1; j-- {
				if m[j][i] == 'Q' {
					continue outer
				}
				floor := len(m) - j - 1
				r, l := i+floor, i-floor
				if r > -1 && r < n && m[j][r] == 'Q' {
					continue outer
				}
				if l > -1 && l < n && m[j][l] == 'Q' {
					continue outer
				}
			}
			mm := append([]string{}, m...)
			mm = append(mm, string(deep))
			f(mm, result)
		}
	}

	result := [][]string{}
	f([]string{}, &result)
	return result
}
