package greedy

import "sort"

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] > people[j][0] {
			return true
		} else if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return false
	})

	for i := 0; i < len(people); i++ {
		for num, cur := people[i][1], i; num < cur; cur-- {
			people[cur-1], people[cur] = people[cur], people[cur-1]
		}
	}
	return people
}
