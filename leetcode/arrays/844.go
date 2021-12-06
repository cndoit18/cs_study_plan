package arrays

// 给定 s 和 t 两个字符串，当它们分别被输入到空白的文本编辑器后，请你判断二者是否相等。# 代表退格字符。
// 如果相等，返回 true ；否则，返回 false 。
func backspaceCompare(s string, t string) bool {
	si, ti := len(s)-1, len(t)-1
	for ; si > -1 || ti > -1; si, ti = si-1, ti-1 {
		for skip := 0; si > -1; {
			if s[si] == '#' {
				skip++
				si--
				continue
			}
			if skip == 0 {
				break
			}
			si--
			skip--
		}
		for skip := 0; ti > -1; {
			if t[ti] == '#' {
				skip++
				ti--
				continue
			}
			if skip == 0 {
				break
			}
			ti--
			skip--
		}
		if ti > -1 && si > -1 && t[ti] != s[si] {
			return false
		}
	}
	return si == ti
}
