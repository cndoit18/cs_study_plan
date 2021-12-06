package arrays

func sortedSquares(nums []int) []int {
	left, right := 0, len(nums)-1
	result := make([]int, len(nums))
	top := len(result) - 1
	for right >= left {
		lsum := nums[left] * nums[left]
		rsum := nums[right] * nums[right]
		if lsum <= rsum {
			right--
			result[top] = rsum
		} else {
			left++
			result[top] = lsum
		}
		top--
	}
	return result
}
