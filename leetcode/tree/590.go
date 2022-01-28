package tree

func postorder(root *Node) []int {
	result := []int{}
	if root == nil {
		return result
	}
	for _, child := range root.Children {
		result = append(result, postorder(child)...)
	}
	result = append(result, root.Val)
	return result
}
