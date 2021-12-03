package arrays

// 给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
// 如果数组中不存在目标值 target，返回 [-1, -1]。
func searchRange(nums []int, target int) []int {
	left, right := 0, len(nums)-1

	for right >= left {
		middle := (right-left)/2 + left
		if nums[middle] == target {
			for left = middle; left > -1 && nums[left] == target; {
				left--
			}
			for right = middle; right < len(nums) && nums[right] == target; {
				right++
			}
			return []int{left + 1, right - 1}
		}
		if nums[middle] > target {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return []int{-1, -1}
}
