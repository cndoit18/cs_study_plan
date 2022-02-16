package tree

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	current := root
	for current != nil {
		if current.Left == nil && current.Val > val {
			current.Left = &TreeNode{
				Val: val,
			}
			break
		}

		if current.Right == nil && current.Val < val {
			current.Right = &TreeNode{
				Val: val,
			}
			break
		}

		if current.Val > val {
			current = current.Left
		} else {
			current = current.Right
		}
	}

	return root
}
