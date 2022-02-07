package tree

import (
	"fmt"
)

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	result := []string{}
	traversal(root, "", &result)
	return result
}

func traversal(node *TreeNode, path string, result *[]string) {
	var current string
	if node == nil {
		return
	}
	if len(path) > 0 {
		current = fmt.Sprint(path, "->", node.Val)
	} else {
		current = fmt.Sprint(node.Val)
	}
	if node.Right == nil && node.Left == nil {
		*result = append(*result, current)
		return
	}
	traversal(node.Right, current, result)

	traversal(node.Left, current, result)
}
