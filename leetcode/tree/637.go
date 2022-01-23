package tree

func averageOfLevels(root *TreeNode) []float64 {
	result := []float64{}
	if root == nil {
		return result
	}
	layer := []*TreeNode{root}
	for len(layer) > 0 {
		tmp := []*TreeNode{}
		sum := 0
		for _, node := range layer {
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			sum += node.Val
		}
		result = append(result, float64(sum)/float64(len(layer)))
		layer = tmp
	}
	return result
}
