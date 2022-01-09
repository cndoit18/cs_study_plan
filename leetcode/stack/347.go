package stack

import "sort"

func topKFrequent(nums []int, k int) []int {
	filter := map[int]int{}
	for _, v := range nums {
		filter[v]++
	}
	r := make([]int, 0, len(filter))

	for k := range filter {
		r = append(r, k)
	}

	sort.Slice(r, func(i, j int) bool {
		return filter[r[i]] > filter[r[j]]
	})
	return r[:k]
}
