package greedy

func partitionLabels(s string) []int {
	result := []int{}
	indexer := [26]int{}
	for i, c := range s {
		indexer[c-'a'] = i
	}
	right, left := 0, -1
	for i := range s {
		right = max(indexer[s[i]-'a'], right)
		if right == i {
			result = append(result, right-left)
			left = right
		}
	}
	return result
}
