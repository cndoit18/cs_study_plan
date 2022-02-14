package tree

func isValidBST(root *TreeNode) bool {
	ll := inorder(root)

	if len(ll) == 0 {
		return false
	}

	for i := 1; i < len(ll); i++ {
		if ll[i] <= ll[i-1] {
			return false
		}
	}

	return true
}

func inorder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := []int{}
	if root.Left != nil {
		result = append(result, inorder(root.Left)...)
	}
	result = append(result, root.Val)
	if root.Right != nil {
		result = append(result, inorder(root.Right)...)
	}

	return result
}
