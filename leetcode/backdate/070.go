package backdate

// 给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
func combine(n int, k int) [][]int {
	var f func([]int, []int, *[][]int)
	nums := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		nums = append(nums, i)
	}
	f = func(ll []int, cur []int, result *[][]int) {
		if k == len(cur) {
			*result = append(*result, cur)
		}
		for i := 0; i < len(ll) && len(cur)+len(ll) >= k; i++ {
			deepcopy := append([]int{}, cur...)
			deepcopy = append(deepcopy, ll[i])
			f(ll[i+1:], deepcopy, result)
		}
	}
	result := [][]int{}
	f(nums, []int{}, &result)
	return result
}
