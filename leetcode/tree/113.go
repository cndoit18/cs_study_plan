package tree

func pathSum(root *TreeNode, targetSum int) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}

	recallPathSumm(root, 0, []int{}, &result, targetSum)
	return result
}

func recallPathSumm(node *TreeNode, sum int, path []int, result *[][]int, tSum int) {
	if node.Right == nil && node.Left == nil && node.Val+sum == tSum {
		*result = append(*result, append(path, node.Val))
		return
	}

	sum += node.Val
	path = append(path, node.Val)
	if node.Right != nil {
		recallPathSumm(node.Right, sum, append([]int{}, path...), result, tSum)
	}
	if node.Left != nil {
		recallPathSumm(node.Left, sum, append([]int{}, path...), result, tSum)
	}
}
