package hashtable

func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	for _, v := range nums1 {
		m[v]++
	}

	result := []int{}
	for _, v := range nums2 {
		if m[v] < 1 {
			delete(m, v)
		}

		if _, ok := m[v]; ok {
			result = append(result, v)
		}
		m[v]--
	}
	return result
}
