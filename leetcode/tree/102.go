package tree

func levelOrder(root *TreeNode) [][]int {
	queue := []*TreeNode{root}
	result := [][]int{}
	for len(queue) > 0 {
		layerNodes := []*TreeNode{}
		layer := []int{}
		for _, node := range queue {
			if node == nil {
				continue
			}
			if node.Left != nil {
				layerNodes = append(layerNodes, node.Left)
			}
			if node.Right != nil {
				layerNodes = append(layerNodes, node.Right)
			}
			layer = append(layer, node.Val)
		}
		queue = layerNodes
		if len(layer) > 0 {
			result = append(result, layer)
		}
	}
	return result
}
