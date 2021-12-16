package hashtable

// 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
func isAnagram(s string, t string) bool {
	m := [26]int{}
	for _, r := range s {
		m[r-rune('a')]++
	}
	for _, r := range t {
		m[r-rune('a')]--
	}

	return m == [26]int{}
}
