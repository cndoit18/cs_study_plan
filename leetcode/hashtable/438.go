package hashtable

// 给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
func findAnagrams(s string, p string) []int {
	m := [26]int{}
	for _, r := range p {
		m[r-'a']++
	}

	slow, fast := 0, len(p)
	result := []int{}
	for fast <= len(s) {
		n := [26]int{}
		for _, r := range s[slow:fast] {
			n[r-'a']++
		}
		if m == n {
			result = append(result, slow)
		}
		slow++
		fast++
	}
	return result
}
