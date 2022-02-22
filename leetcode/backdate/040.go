package backdate

import (
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var f func([]int, []int, int, *[][]int)
	f = func(ll []int, cur []int, sum int, result *[][]int) {
		if sum == target {
			*result = append(*result, cur)
			return
		}

		for i := 0; i < len(ll) && sum+ll[i] <= target; i++ {
			if i > 0 && ll[i] == ll[i-1] {
				continue
			}
			deepcopy := append([]int{}, cur...)
			deepcopy = append(deepcopy, ll[i])
			f(ll[i+1:], deepcopy, sum+ll[i], result)
		}
	}
	result := [][]int{}
	f(candidates, []int{}, 0, &result)
	return result
}
