package greedy

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	min := func(x, y int) int {
		if x > y {
			return y
		}

		return x
	}
	max := func(x, y int) int {
		if x > y {
			return x
		}

		return y
	}
	var per [][]int
	for i := 1; i < len(intervals); {
		if intervals[i][0] <= intervals[i-1][1] {
			if i > 1 {
				per = intervals[:i-1]
			} else {
				per = [][]int{}
			}
			intervals = append(per, append([][]int{{min(intervals[i-1][0], intervals[i][0]), max(intervals[i-1][1], intervals[i][1])}}, intervals[i+1:]...)...)
			continue
		}
		i++
	}
	return intervals
}
