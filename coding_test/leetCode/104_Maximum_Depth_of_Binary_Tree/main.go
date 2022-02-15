package main

import (
	"fmt"
	"math"
)

// Given the root of a binary tree, return its maximum depth.

// A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

// Example 1:

//        3
//      /   \
//     9     20
//         /    \
//        15     7

// Input: root = [3,9,20,null,null,15,7]
// Output: 3
// Example 2:

// Input: root = [1,null,2]
// Output: 2

// Constraints:

// The number of nodes in the tree is in the range [0, 104].
// -100 <= Node.val <= 100

//Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
 
func maxDepth(root *TreeNode) int {
    if root == nil {	
		return 0
	}

    return int(getDepth(root, 1))
}

func getDepth(root *TreeNode, parentDepth float64) float64 {
    leftDepth := parentDepth
	rightDepth := parentDepth

	if root.Left != nil {
		leftDepth++
		leftDepth = getDepth(root.Left, leftDepth)
	}

	if root.Right != nil {
		rightDepth++
		rightDepth = getDepth(root.Right, rightDepth)
	}
    
    return math.Max(leftDepth, rightDepth)
}

func main() {
	node0 := TreeNode { Val: 3 }
	node1 := TreeNode { Val: 9 }
	node2 := TreeNode { Val: 20 }
	node6 := TreeNode { Val: 15 }
	node7 := TreeNode { Val: 7 }

	node2.Left = &node6
	node2.Right = &node7

	node0.Left = &node1
	node0.Right = &node2

	fmt.Println(maxDepth(&node0))
}