package tree

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	r, l := maxDepth(root.Right), maxDepth(root.Left)
	if r > l {
		return r + 1
	}
	return l + 1
}
