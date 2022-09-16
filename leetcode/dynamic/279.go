package dynamic

import "math"

func numSquares(n int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i < len(dp); i++ {
		dp[i] = math.MaxInt - 1
	}
	for i := 2; i <= n; i++ {
		for j := i; j > 0; j-- {
			if i >= j*j {
				dp[i] = min(dp[i-j*j]+1, dp[i])
			}
		}
	}

	return dp[n]
}
