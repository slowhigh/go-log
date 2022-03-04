package maximum_depth_of_binary_tree_104

import "testing"

func Test_maxDepth(t *testing.T) {
	node0 := TreeNode{Val: 3}
	node1 := TreeNode{Val: 9}
	node2 := TreeNode{Val: 20}
	node6 := TreeNode{Val: 15}
	node7 := TreeNode{Val: 7}

	node2.Left = &node6
	node2.Right = &node7

	node0.Left = &node1
	node0.Right = &node2

	if maxDepth(&node0) != 3 {
		t.Error("Wrong output")
	}
}
