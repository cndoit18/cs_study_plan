package tree

func connect(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	level := []*TreeNode{root}
	for len(level) > 0 {
		ll := []*TreeNode{}
		if level[0].Left != nil {
			ll = append(ll, level[0].Left)
		}
		if level[0].Right != nil {
			ll = append(ll, level[0].Right)
		}
		for i := 0; i < len(level)-1; i++ {
			level[i].Next = level[i+1]
			if level[i+1].Left != nil {
				ll = append(ll, level[i+1].Left)
			}
			if level[i+1].Right != nil {
				ll = append(ll, level[i+1].Right)
			}
		}
		level = ll
	}
	return root
}
