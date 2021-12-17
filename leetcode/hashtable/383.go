package hashtable

// 给你两个字符串：ransomNote 和 magazine ，判断 ransomNote 能不能由 magazine 里面的字符构成。
// 如果可以，返回 true ；否则返回 false 。
// magazine 中的每个字符只能在 ransomNote 中使用一次。

func canConstruct(ransomNote string, magazine string) bool {
	m := [26]int{}
	for _, r := range ransomNote {
		m[r-'a']--
	}
	for _, r := range magazine {
		if m[r-'a'] > -1 {
			continue
		}
		m[r-'a']++
	}

	return m == [26]int{}
}
