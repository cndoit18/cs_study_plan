package dynamic

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	f := func(e, o int) int {
		if o == 1 {
			return 0
		}
		return e
	}
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		if i > 0 && dp[i-1][0] == 0 {
			dp[i][0] = 0
			continue
		}
		dp[i][0] = f(1, obstacleGrid[i][0])
	}
	for j := 0; j < n; j++ {
		if j > 0 && dp[0][j-1] == 0 {
			dp[0][j] = 0
			continue
		}
		dp[0][j] = f(1, obstacleGrid[0][j])
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = f(dp[i][j-1]+dp[i-1][j], obstacleGrid[i][j])
		}
	}
	return dp[m-1][n-1]
}
