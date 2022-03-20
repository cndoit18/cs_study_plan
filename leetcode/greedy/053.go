package greedy

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sum := nums[0]
	max := sum
	for j := 1; j < len(nums); j++ {
		if sum+nums[j] > nums[j] {
			sum += nums[j]
		} else {
			sum = nums[j]
		}

		if max < sum {
			max = sum
		}
	}

	return max
}
