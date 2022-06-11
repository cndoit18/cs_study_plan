package dynamic

func lastStoneWeightII(stones []int) int {
	weight := 0
	for _, w := range stones {
		weight += w
	}

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	target := int(weight / 2)
	dp := make([]int, weight)
	for _, w := range stones {
		for j := target; j >= w; j-- {
			dp[j] = max(dp[j], dp[j-w]+w)
		}
	}
	return weight - (dp[target] << 1)
}
