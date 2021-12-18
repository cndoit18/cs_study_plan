package hashtable

// 给定两个数组，编写一个函数来计算它们的交集。
func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]struct{})
	for _, v := range nums1 {
		m[v] = struct{}{}
	}

	result := []int{}
	for _, v := range nums2 {
		if _, ok := m[v]; ok {
			result = append(result, v)
		}
		delete(m, v)
	}
	return result
}
