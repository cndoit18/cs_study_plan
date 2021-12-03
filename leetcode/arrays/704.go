package arrays

// 给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for right >= left {
		middle := (right-left)/2 + left
		if nums[middle] == target {
			return middle
		}

		if nums[middle] > target {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return -1
}
