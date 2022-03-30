package greedy

func candy(ratings []int) int {
	candies := make([]int, len(ratings))
	for i := range candies {
		candies[i] = 1
	}

	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	for i := len(ratings) - 2; i > -1; i-- {
		if ratings[i] > ratings[i+1] {
			candies[i] = max(candies[i], candies[i+1]+1)
		}
	}

	sum := 0
	for _, v := range candies {
		sum += v
	}
	return sum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
