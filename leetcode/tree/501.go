package tree

func findMode(root *TreeNode) []int {
	var (
		pre      *TreeNode
		maxCount int
		count    int
		histroy  map[int]struct{}
		travel   func(*TreeNode)
	)

	travel = func(current *TreeNode) {
		if current == nil {
			return
		}

		travel(current.Left)

		if pre != nil && pre.Val != current.Val {
			count = 1
		}

		if pre == nil || pre.Val == current.Val {
			count++
		}

		if count > maxCount {
			maxCount = count
			histroy = make(map[int]struct{})
		}
		if count == maxCount {
			histroy[current.Val] = struct{}{}
		}

		pre = current
		travel(current.Right)
	}

	travel(root)
	result := make([]int, 0, len(histroy))
	for k := range histroy {
		result = append(result, k)
	}
	return result
}
