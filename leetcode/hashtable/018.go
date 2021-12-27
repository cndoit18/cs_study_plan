package hashtable

import "sort"

// 给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的四元组 [nums[a], nums[b], nums[c], nums[d]] （若两个四元组元素一一对应，则认为两个四元组重复）：
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	for i := 0; i < len(nums); i++ {

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			left, right := j+1, len(nums)-1
			for left < right {
				if nums[i]+nums[j]+nums[left]+nums[right] > target {
					right--
					continue
				}
				if nums[i]+nums[j]+nums[left]+nums[right] == target {
					if left > j+1 && nums[left] == nums[left-1] {
						left++
						continue
					}
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
				}
				left++
			}
		}
	}
	return result
}
