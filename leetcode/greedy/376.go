package greedy

func wiggleMaxLength(nums []int) int {
	preDiff := 0
	sum := 1
	for r := 1; r < len(nums); r++ {
		diff := nums[r] - nums[r-1]

		if diff == 0 || (preDiff > 0 && diff > 0 || (preDiff < 0 && diff < 0)) {
			continue
		}
		sum++
		preDiff = diff
	}
	return sum
}
