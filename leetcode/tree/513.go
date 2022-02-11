package tree

func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var sum int
	layer := []*TreeNode{root}
	for len(layer) != 0 {
		sum = layer[0].Val
		ll := []*TreeNode{}
		for i := range layer {
			node := layer[i]
			if node.Left != nil {
				ll = append(ll, node.Left)
			}
			if node.Right != nil {
				ll = append(ll, node.Right)
			}
		}
		layer = ll
	}

	return sum
}
