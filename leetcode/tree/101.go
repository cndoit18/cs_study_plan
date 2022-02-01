package tree

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return compare(root.Right, root.Left)
}

func compare(left, right *TreeNode) bool {
	if left == nil || right == nil {
		return left == right
	}

	if left.Val != right.Val {
		return false
	}

	return compare(left.Left, right.Right) && compare(left.Right, right.Left)
}
