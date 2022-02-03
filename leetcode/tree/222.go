package tree

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return countNodes(root.Right) + countNodes(root.Left) + 1
}
