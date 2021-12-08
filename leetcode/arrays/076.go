package arrays

// 给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
func minWindow(s string, t string) string {
	tm := map[byte]int{}
	check := map[byte]struct{}{}
	for _, c := range []byte(t) {
		tm[c]++
		check[c] = struct{}{}
	}
	slow, fast, minStr := 0, 0, ""
	for fast < len(s) {
		key := s[fast]
		if nums, ok := tm[key]; ok {
			nums--
			tm[key] = nums
			if nums < 1 {
				delete(check, key)
			}
		}
		fast++

		for len(check) == 0 {
			tmp := s[slow:fast]
			if len(tmp) < len(minStr) || len(minStr) == 0 {
				minStr = tmp
			}
			key := s[slow]
			if nums, ok := tm[key]; ok {
				nums++
				tm[key] = nums
				if nums > 0 {
					check[key] = struct{}{}
				}

			}
			slow++
		}
	}
	return minStr
}
