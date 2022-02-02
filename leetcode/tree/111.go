package tree

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	r := minDepth(root.Right)
	l := minDepth(root.Left)

	if root.Left == nil && root.Right != nil {
		return 1 + r
	}

	if root.Right == nil && root.Left != nil {
		return 1 + l
	}

	if r > l {
		return 1 + l
	}

	return 1 + r
}
