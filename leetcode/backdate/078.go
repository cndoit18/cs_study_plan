package backdate

func subsets(nums []int) [][]int {
	var f func(cur []int, rest []int, result *[][]int)
	f = func(cur, rest []int, result *[][]int) {
		*result = append(*result, cur)
		for i := 0; i < len(rest); i++ {
			tmp := append([]int{}, cur...)
			tmp = append(tmp, rest[i])
			f(tmp, rest[i+1:], result)
		}
	}
	result := [][]int{}
	f([]int{}, nums, &result)
	return result
}
