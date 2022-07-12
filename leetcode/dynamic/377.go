package dynamic

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1

	for j := 0; j <= target; j++ {
		for _, num := range nums {
			if j < num {
				continue
			}
			dp[j] += dp[j-num]
		}
	}

	return dp[target]
}
