package tree

func getMinimumDifference(root *TreeNode) int {
	ll := inorder(root)
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	if len(ll) <= 1 {
		return 0
	}

	r := abs(ll[1] - ll[0])
	for i := 1; i < len(ll); i++ {
		r = min(abs(ll[i]-ll[i-1]), r)
	}
	return r
}
