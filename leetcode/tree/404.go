package tree

func sumOfLeftLeaves(root *TreeNode) int {
	sum := 0
	if root.Left != nil {
		if root.Left.Right == nil && root.Left.Left == nil {
			sum += root.Left.Val
		} else {
			sum += sumOfLeftLeaves(root.Left)
		}
	}
	if root.Right != nil {
		sum += sumOfLeftLeaves(root.Right)
	}

	return sum
}
