package hashtable

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	result := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			continue
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1
		for left < right {
			if nums[i]+nums[left]+nums[right] > 0 {
				right--
				continue
			}
			if nums[i]+nums[left]+nums[right] == 0 {
				if left > i+1 && nums[left] == nums[left-1] {
					left++
					continue
				}
				result = append(result, []int{nums[i], nums[left], nums[right]})
			}
			left++
		}
	}
	return result
}
