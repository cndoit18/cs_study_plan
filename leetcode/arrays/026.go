package arrays

// 给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。
// 不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
func removeDuplicates(nums []int) int {
	slow, fast := 0, 0
	for fast < len(nums) {
		for slow == fast || (fast < len(nums) && nums[fast] == nums[slow]) {
			fast++
		}
		slow++
		if fast >= len(nums) {
			break
		}
		nums[slow] = nums[fast]
	}
	return slow
}
