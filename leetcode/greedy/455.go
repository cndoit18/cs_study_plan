package greedy

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Slice(g, func(i, j int) bool {
		return g[i] < g[j]
	})
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})

	sum := 0
	for i, j := len(g)-1, len(s)-1; i > -1 && j > -1; {
		if g[i] > s[j] {
			i--
			continue
		}
		j, i, sum = j-1, i-1, sum+1
	}
	return sum
}
