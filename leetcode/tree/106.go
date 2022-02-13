package tree

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 && len(postorder) == 0 {
		return nil
	}

	root := &TreeNode{}
	root.Val = postorder[len(postorder)-1]
	index := 0
	for index < len(inorder) && inorder[index] != root.Val {
		index++
	}
	root.Left = buildTree(inorder[:index], postorder[:index])
	if index <= len(postorder)-1 {
		root.Right = buildTree(inorder[index+1:], postorder[index:len(postorder)-1])
	}

	return root
}
