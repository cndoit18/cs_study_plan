package dynamic

func rob(nums []int) int {
	dp := make([]int, len(nums)+1)
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	for i := 1; i < len(dp); i++ {
		if i < 2 {
			dp[i] = max(dp[i-1], nums[i-1])
			continue
		}

		dp[i] = max(dp[i-1], dp[i-2]+nums[i-1])
	}
	return dp[len(dp)-1]
}
