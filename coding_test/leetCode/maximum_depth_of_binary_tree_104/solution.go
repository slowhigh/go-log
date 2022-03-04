package maximum_depth_of_binary_tree_104

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left) + 1
	rightDepth := maxDepth(root.Right) + 1

	if leftDepth > rightDepth {
		return leftDepth
	} else {
		return rightDepth
	}
}
