package backdate

func combinationSum3(k int, n int) [][]int {
	var f func([]int, int, map[int]struct{}, *[][]int)
	l := []int{}
	for i := 1; i < 10; i++ {
		l = append(l, i)
	}
	f = func(l []int, sum int, m map[int]struct{}, result *[][]int) {
		if len(m) == k {
			if sum == n {
				layer := []int{}
				for k := range m {
					layer = append(layer, k)
				}
				*result = append(*result, layer)
			}
			return
		}
		for i, v := range l {
			if _, ok := m[v]; !ok && sum+v <= n && len(m) < k {
				deepcopy := map[int]struct{}{}
				for k := range m {
					deepcopy[k] = struct{}{}
				}
				deepcopy[v] = struct{}{}
				f(l[i+1:], sum+v, deepcopy, result)
			}
		}
	}
	result := [][]int{}
	f(l, 0, map[int]struct{}{}, &result)
	return result
}
