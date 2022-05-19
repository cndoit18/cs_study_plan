package greedy

func maxProfit_714(prices []int, fee int) int {
	sum := 0
	buy := prices[0] + fee
	for i := 1; i < len(prices); i++ {
		if prices[i]+fee < buy {
			buy = prices[i] + fee
		} else if prices[i] > buy {
			sum += (prices[i] - buy)
			buy = prices[i]
		}
	}
	return sum
}
