package tree

func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	return sumBST(root, &sum)
}

func sumBST(root *TreeNode, sum *int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		*sum += root.Val
		root.Val = *sum
		return root
	}

	sumBST(root.Right, sum)

	*sum += root.Val
	root.Val = *sum

	sumBST(root.Left, sum)
	return root
}
