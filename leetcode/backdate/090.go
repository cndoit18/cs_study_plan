package backdate

import "sort"

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	var f func(cur []int, rest []int, result *[][]int)
	f = func(cur, rest []int, result *[][]int) {
		*result = append(*result, cur)
		filter := map[int]struct{}{}
		for i := 0; i < len(rest); i++ {
			if _, ok := filter[rest[i]]; ok {
				continue
			}
			filter[rest[i]] = struct{}{}
			tmp := append([]int{}, cur...)
			tmp = append(tmp, rest[i])
			f(tmp, rest[i+1:], result)
		}
	}
	result := [][]int{}
	f([]int{}, nums, &result)
	return result
}
