package greedy

import (
	"math"
	"sort"
)

func largestSumAfterKNegations(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return math.Abs(float64(nums[i])) < math.Abs(float64(nums[j]))
	})

	for i := len(nums) - 1; i > -1 && k > 0; i-- {
		if nums[i] < 0 {
			k--
			nums[i] = -nums[i]
		}
	}

	if k%2 != 0 {
		nums[0] = -nums[0]
	}

	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}
