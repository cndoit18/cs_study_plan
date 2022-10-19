package dynamic

func rob_213(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	rob := func(nums []int) int {
		dp := make([]int, len(nums)+1)
		for i := 1; i < len(dp); i++ {
			if i < 2 {
				dp[i] = max(dp[i-1], nums[i-1])
				continue
			}

			dp[i] = max(dp[i-1], dp[i-2]+nums[i-1])
		}
		return dp[len(dp)-1]
	}
	return max(rob(nums[:len(nums)-1]), rob(nums[1:]))
}
