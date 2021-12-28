package string

func reverseString(s []byte) {
	left, right := 0, len(s)-1
	for ; right > left; right, left = right-1, left+1 {
		s[left], s[right] = s[right], s[left]
	}
}
