package string

// 给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。如果不存在，则返回  -1 。

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	next := make([]int, len(needle))
	for current, i := 1, 0; current < len(next); current++ {
		for i > 0 && needle[i] != needle[current] {
			i = next[i-1]
		}
		if needle[i] == needle[current] {
			i++
		}
		next[current] = i
	}

	for current, j := 0, 0; current < len(haystack); {
		for ; current < len(haystack) && haystack[current] == needle[j]; j, current = j+1, current+1 {
			if j == len(needle)-1 {
				return current - len(needle) + 1
			}
		}
		if j > 0 {
			j = next[j-1]
			continue
		}
		current++
	}

	return -1
}
