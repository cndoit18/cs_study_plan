package tree

func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	level := []*TreeNode{root}
	sum := []int{}
	for len(level) > 0 {
		ll := []*TreeNode{}
		max := level[0].Val
		for _, node := range level {
			if node.Val > max {
				max = node.Val
			}
			if node.Left != nil {
				ll = append(ll, node.Left)
			}
			if node.Right != nil {
				ll = append(ll, node.Right)
			}
		}
		sum = append(sum, max)
		level = ll
	}
	return sum
}
