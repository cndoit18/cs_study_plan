package backdate

func permute(nums []int) [][]int {
	var f func(set []int, rest []int, result *[][]int)
	f = func(set []int, rest []int, result *[][]int) {
		if len(set) == 0 {
			*result = append(*result, rest)
			return
		}
		for i := range set {
			subset := append([]int{}, set[:i]...)
			subset = append(subset, set[i+1:]...)

			f(subset, append(append([]int{}, rest...), set[i]), result)
		}
	}
	result := [][]int{}
	f(nums, []int{}, &result)
	return result
}
