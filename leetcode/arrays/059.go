package arrays

// 给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。
func generateMatrix(n int) [][]int {
	top, bottom := 0, n-1
	left, right := 0, n-1
	nums := 1
	tar := n * n
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	for nums <= tar {
		for i := left; i <= right; i++ {
			matrix[top][i] = nums
			nums++
		}
		top++
		for i := top; i <= bottom; i++ {
			matrix[i][right] = nums
			nums++
		}
		right--
		for i := right; i >= left; i-- {
			matrix[bottom][i] = nums
			nums++
		}
		bottom--
		for i := bottom; i >= top; i-- {
			matrix[i][left] = nums
			nums++
		}
		left++
	}
	return matrix
}
