package dynamic

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	f := func(s string) (int, int) {
		zn, on := 0, 0
		for _, v := range s {
			if v == rune('1') {
				on++
				continue
			}
			zn++
		}
		return zn, on
	}

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	for i := range strs {
		s := strs[i]
		zn, on := f(s)
		for x := m; x >= zn; x-- {
			for y := n; y >= on; y-- {
				dp[x][y] = max(dp[x][y], dp[x-zn][y-on]+1)
			}
		}
	}
	return dp[m][n]
}
