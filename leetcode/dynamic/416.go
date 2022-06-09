package dynamic

func canPartition(nums []int) bool {
	sum := 0
	for _, n := range nums {
		sum += n
	}

	if sum%2 == 1 {
		return false
	}

	target := sum / 2
	dp := make([]int, sum)
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	for _, num := range nums {
		for j := target; j >= num; j-- {
			dp[j] = max(dp[j], dp[j-num]+num)
		}
	}

	return dp[target] == target
}
