package tree

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == key {
		if root.Right == nil {
			return root.Left
		}

		current := root.Right
		for current.Left != nil {
			current = current.Left
		}
		current.Left = root.Left
		return root.Right
	}

	root.Left = deleteNode(root.Left, key)
	root.Right = deleteNode(root.Right, key)
	return root
}
