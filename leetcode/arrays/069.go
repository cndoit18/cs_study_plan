package arrays

// 给你一个非负整数 x ，计算并返回 x 的 算术平方根 。
// 由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。
func mySqrt(x int) int {
	left, right, middle := 0, x, 0

	for right >= left {
		middle = (right-left)/2 + left
		sum := middle * middle
		if sum == x {
			return middle
		}
		if sum > x {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}

	if middle*middle > x {
		return middle - 1
	}
	return middle
}
