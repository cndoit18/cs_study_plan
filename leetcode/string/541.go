package string

func reverseStr(s string, k int) string {
	result := []byte(s)
	for i := 0; i < len(s); i += k << 1 {
		left := i
		right := i + k - 1
		if right >= len(s) {
			right = len(s) - 1
		}
		for ; right > left; right, left = right-1, left+1 {
			result[left], result[right] = result[right], result[left]
		}
	}
	return string(result)
}
