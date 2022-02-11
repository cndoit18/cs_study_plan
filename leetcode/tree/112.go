package tree

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	return recallHasPathSum(root, 0, targetSum)
}

func recallHasPathSum(node *TreeNode, sum int, tSum int) bool {
	if node.Right == nil && node.Left == nil {
		return node.Val+sum == tSum
	}
	result := false
	sum += node.Val
	if node.Right != nil {
		result = result || recallHasPathSum(node.Right, sum, tSum)
	}
	if node.Left != nil {
		result = result || recallHasPathSum(node.Left, sum, tSum)
	}

	return result
}
