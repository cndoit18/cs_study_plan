package greedy

import "sort"

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	arrow := 1
	preEnd := points[0][1]
	for i := 1; i < len(points); i++ {
		if points[i][0] > preEnd {
			arrow++
			preEnd = points[i][1]
			continue
		}

		if points[i][1] < preEnd {
			preEnd = points[i][1]
		}
	}
	return arrow
}
