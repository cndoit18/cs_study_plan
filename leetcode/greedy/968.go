package greedy

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func minCameraCover(root *TreeNode) int {
	result := 0
	var dst func(node *TreeNode) int
	dst = func(node *TreeNode) int {
		if node == nil {
			return 1
		}
		l, r := dst(node.Left), dst(node.Right)
		if l == 0 || r == 0 {
			result++
			return 2
		}
		if l == 2 || r == 2 {
			return 1
		}

		return 0
	}

	if dst(root) == 0 {
		result++
	}
	return result
}
