package backdate

import "sort"

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	var f func(set []int, rest []int, result *[][]int)
	f = func(set []int, rest []int, result *[][]int) {
		if len(set) == 0 {
			*result = append(*result, rest)
			return
		}
		for i := range set {
			if i > 0 && set[i] == set[i-1] {
				continue
			}
			subset := append([]int{}, set[:i]...)
			subset = append(subset, set[i+1:]...)

			f(subset, append(append([]int{}, rest...), set[i]), result)
		}
	}
	result := [][]int{}
	f(nums, []int{}, &result)
	return result
}
