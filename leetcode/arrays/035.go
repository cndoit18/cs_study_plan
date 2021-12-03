package arrays

// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
func searchInsert(nums []int, target int) int {
	left, right, middle := 0, len(nums)-1, 0

	if len(nums) == 0 {
		return 0
	}

	for right >= left {
		middle = (right-left)/2 + left

		if nums[middle] == target {
			return middle
		}

		if nums[middle] > target {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}

	if nums[middle] > target {
		return middle
	}
	return middle + 1
}
