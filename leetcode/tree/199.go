package tree

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	arrays := []*TreeNode{root}
	result := []int{}
	for len(arrays) != 0 {
		layer := []*TreeNode{}
		for _, node := range arrays {
			if node.Left != nil {
				layer = append(layer, node.Left)
			}
			if node.Right != nil {
				layer = append(layer, node.Right)
			}
		}
		if len(arrays) > 0 {
			result = append(result, arrays[len(arrays)-1].Val)
		}
		arrays = layer
	}

	return result
}
