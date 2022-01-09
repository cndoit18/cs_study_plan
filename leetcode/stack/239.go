package stack

import "container/list"

func maxSlidingWindow(nums []int, k int) []int {
	ls := list.New()
	result := make([]int, 0, len(nums))
	for left, right, i := 0, k, 0; right <= len(nums); {
		for ; i < len(nums) && i < right; i++ {

			for ls.Len() != 0 && ls.Back().Value.(int) < nums[i] {
				ls.Remove(ls.Back())
			}
			ls.PushBack(nums[i])
		}
		result = append(result, ls.Front().Value.(int))
		if ls.Len() > 0 && ls.Front().Value.(int) == nums[left] {
			ls.Remove(ls.Front())
		}
		left, right = left+1, right+1
	}
	return result
}
