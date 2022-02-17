package tree

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val < low || root.Val > high {
		if root.Right == nil {
			return trimBST(root.Left, low, high)
		}

		current := root.Right
		for current.Left != nil {
			current = current.Left
		}
		current.Left = root.Left
		return trimBST(root.Right, low, high)
	}

	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root
}
