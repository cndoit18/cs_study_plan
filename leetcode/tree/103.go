package tree

func zigzagLevelOrder(root *TreeNode) [][]int {
	queue := []*TreeNode{root}
	result := [][]int{}
	rev := false
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
		if rev {
			for i, j := 0, len(layer)-1; i < j; i, j = i+1, j-1 {
				layer[i], layer[j] = layer[j], layer[i]
			}
		}
		rev = !rev
		if len(layer) > 0 {
			result = append(result, layer)
		}
	}
	return result
}
