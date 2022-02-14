package tree

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	var (
		val   int
		left  *TreeNode
		right *TreeNode
	)

	if root1 == nil && root2 == nil {
		return nil
	}

	if root1 != nil && root2 != nil {
		return &TreeNode{
			Val:   root1.Val + root2.Val,
			Left:  mergeTrees(root1.Left, root2.Left),
			Right: mergeTrees(root1.Right, root2.Right),
		}
	}

	if root1 != nil {
		val = root1.Val
		left = mergeTrees(root1.Left, nil)
		right = mergeTrees(root1.Right, nil)
	} else {
		val = root2.Val
		left = mergeTrees(root2.Left, nil)
		right = mergeTrees(root2.Right, nil)
	}

	return &TreeNode{
		Val:   val,
		Left:  left,
		Right: right,
	}
}
