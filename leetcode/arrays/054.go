package arrays

// 给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
func spiralOrder(matrix [][]int) []int {
	top, bottom := 0, len(matrix)-1
	if len(matrix) == 0 {
		return []int{}
	}
	left, right := 0, len(matrix[0])-1
	nums := 0
	tar := (bottom + 1) * (right + 1)
	result := make([]int, tar)
	for nums < tar {
		for i := left; i <= right && nums < tar; i++ {
			result[nums] = matrix[top][i]
			nums++
		}
		top++
		for i := top; i <= bottom && nums < tar; i++ {
			result[nums] = matrix[i][right]
			nums++
		}
		right--
		for i := right; i >= left && nums < tar; i-- {
			result[nums] = matrix[bottom][i]
			nums++
		}
		bottom--
		for i := bottom; i >= top && nums < tar; i-- {
			result[nums] = matrix[i][left]
			nums++
		}
		left++
	}
	return result
}
