package hashtable

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, v := range nums {
		sum := target - v
		if n, ok := m[sum]; ok {
			return []int{n, i}
		}
		m[v] = i
	}
	return []int{}
}
