package dynamic

import "math"

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}

	if int(math.Abs(float64(target))) > sum {
		return 0
	}
	if (sum+target)%2 == 1 {
		return 0
	}

	w := (sum + target) / 2
	dp := make([]int, sum+1)
	dp[0] = 1
	for _, num := range nums {
		for j := w; j >= num; j-- {
			dp[j] += dp[j-num]
		}
	}
	return dp[w]
}
