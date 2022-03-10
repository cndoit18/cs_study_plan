package backdate

func findSubsequences(nums []int) [][]int {
	var f func(subset []int, rest []int, result *[][]int)
	f = func(subset []int, rest []int, result *[][]int) {
		if len(subset) > 1 {
			*result = append(*result, subset)
		}
		var last int
		cache := map[int]struct{}{}
		if len(subset) > 0 {
			last = subset[len(subset)-1]
		}
		for i := range rest {
			if len(subset) == 0 || rest[i] >= last {
				if _, ok := cache[rest[i]]; ok {
					continue
				}
				cache[rest[i]] = struct{}{}
				f(append(append([]int{}, subset...), rest[i]), rest[i+1:], result)
			}
		}
	}
	result := [][]int{}
	f([]int{}, nums, &result)
	return result
}
