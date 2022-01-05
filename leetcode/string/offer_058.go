package string

// 字符串的左旋转操作是把字符串前面的若干个字符转移到字符串的尾部。请定义一个函数实现字符串左旋转操作的功能。
// 比如，输入字符串"abcdefg"和数字2，该函数将返回左旋转两位得到的结果"cdefgab"。
func reverseLeftWords(s string, n int) string {
	arrays := []byte(s)
	for font, end := 0, n-1; font < end; font, end = font+1, end-1 {
		arrays[font], arrays[end] = arrays[end], arrays[font]
	}
	for font, end := n, len(arrays)-1; font < end; font, end = font+1, end-1 {
		arrays[font], arrays[end] = arrays[end], arrays[font]
	}

	for font, end := 0, len(arrays)-1; font < end; font, end = font+1, end-1 {
		arrays[font], arrays[end] = arrays[end], arrays[font]
	}
	return string(arrays)
}
