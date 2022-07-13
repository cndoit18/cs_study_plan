package dynamic

const MAX_INT = int(^uint(0) >> 1)

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = MAX_INT
	}
	min := func(x, y int) int {
		if x > y && y > 0 {
			return y
		}
		return x
	}
	for i := 0; i <= amount; i++ {
		for _, coin := range coins {
			if i < coin || (dp[i] == MAX_INT && dp[i-coin] == MAX_INT) {
				continue
			}

			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}

	if dp[amount] == MAX_INT {
		return -1
	}
	return dp[amount]
}
