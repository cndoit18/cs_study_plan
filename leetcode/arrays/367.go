package arrays

// 给定一个 正整数 num ，编写一个函数，如果 num 是一个完全平方数，则返回 true ，否则返回 false 。
func isPerfectSquare(num int) bool {
	left, right := 0, num
	for right >= left {
		middle := (right-left)/2 + left
		sum := middle * middle
		if sum == num {
			return true
		}
		if sum > num {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return false
}
