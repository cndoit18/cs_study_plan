package tree

func levelOrderNode(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}
	sum := append([][]int{}, []int{root.Val})
	layer := root.Children
	for len(layer) > 0 {
		tmp := []*Node{}
		ll := make([]int, 0, len(layer))
		for _, node := range layer {
			ll = append(ll, node.Val)
			tmp = append(tmp, node.Children...)
		}
		sum = append(sum, ll)
		layer = tmp
	}
	return sum
}
