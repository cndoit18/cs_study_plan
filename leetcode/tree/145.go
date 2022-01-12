package tree

func postorderTraversal(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	result = append(result, postorderTraversal(root.Right)...)
	result = append(result, postorderTraversal(root.Left)...)
	result = append(result, root.Val)

	return result
}
