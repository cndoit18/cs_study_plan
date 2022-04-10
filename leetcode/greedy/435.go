package greedy

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	left, nums := 0, 0
	for i := range intervals {
		if i == 0 {
			left = intervals[i][1]
			continue
		}
		if intervals[i][0] < left && intervals[i][1] >= left {
			nums++
			continue
		}
		left = intervals[i][1]
	}
	return nums
}
