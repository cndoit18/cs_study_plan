package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
	Next  *TreeNode
}

type Node struct {
	Val      int
	Children []*Node
}
