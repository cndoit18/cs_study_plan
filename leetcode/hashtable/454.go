package hashtable

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	m := make(map[int]int)
	count := 0
	for _, x := range nums1 {
		for _, y := range nums2 {
			m[x+y]++
		}
	}
	for _, x := range nums3 {
		for _, y := range nums4 {
			if v, ok := m[0-(x+y)]; ok {
				count += v
			}
		}
	}
	return count
}
