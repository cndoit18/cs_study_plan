package dynamic

func integerBreak(n int) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i < n+1; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(j*(i-j), max(dp[j]*(i-j), dp[i]))
		}
	}
	return dp[n]
}
