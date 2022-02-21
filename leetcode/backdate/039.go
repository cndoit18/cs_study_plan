package backdate

import "sort"

// candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var f func([]int, []int, int, *[][]int)
	f = func(ll []int, cur []int, sum int, result *[][]int) {
		if sum == target {
			*result = append(*result, cur)
			return
		}

		for i := 0; i < len(ll) && sum+ll[i] <= target; i++ {
			deepcopy := append([]int{}, cur...)
			deepcopy = append(deepcopy, ll[i])
			f(ll[i:], deepcopy, sum+ll[i], result)
		}
	}
	result := [][]int{}
	f(candidates, []int{}, 0, &result)
	return result
}
