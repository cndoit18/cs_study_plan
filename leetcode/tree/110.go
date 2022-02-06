package tree

// 给定一个二叉树，判断它是否是高度平衡的二叉树

func isBalanced(root *TreeNode) bool {
	return getHeight(root) != -1
}

func getHeight(root *TreeNode) int {
	var r, l int
	if root == nil {
		return 0
	}

	if r = getHeight(root.Right); r == -1 {
		return -1
	}
	if l = getHeight(root.Left); l == -1 {
		return -1
	}

	if r-l > 1 || r-l < -1 {
		return -1
	}

	if r > l {
		return r + 1
	}
	return l + 1
}
