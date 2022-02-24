package main

// Given a reference of a node in a connected undirected graph.

// Return a deep copy (clone) of the graph.

// Each node in the graph contains a value (int) and a list (List[Node]) of its neighbors.

// class Node {
//     public int val;
//     public List<Node> neighbors;
// }

// Test case format:

// For simplicity, each node's value is the same as the node's index (1-indexed). For example, the first node with val == 1, the second node with val == 2, and so on. The graph is represented in the test case using an adjacency list.

// An adjacency list is a collection of unordered lists used to represent a finite graph. Each list describes the set of neighbors of a node in the graph.

// The given node will always be the first node with val = 1. You must return the copy of the given node as a reference to the cloned graph.

// Example 1:
//    1 -- 2
//    |    |
//    4 -- 3

// Input: adjList = [[2,4],[1,3],[2,4],[1,3]]
// Output: [[2,4],[1,3],[2,4],[1,3]]
// Explanation: There are 4 nodes in the graph.
// 1st node (val = 1)'s neighbors are 2nd node (val = 2) and 4th node (val = 4).
// 2nd node (val = 2)'s neighbors are 1st node (val = 1) and 3rd node (val = 3).
// 3rd node (val = 3)'s neighbors are 2nd node (val = 2) and 4th node (val = 4).
// 4th node (val = 4)'s neighbors are 1st node (val = 1) and 3rd node (val = 3).

// Example 2:
// Input: adjList = [[]]
// Output: [[]]
// Explanation: Note that the input contains one empty list. The graph consists of only one node with val = 1 and it does not have any neighbors.

// Example 3:
// Input: adjList = []
// Output: []
// Explanation: This an empty graph, it does not have any nodes.

// Constraints:
// The number of nodes in the graph is in the range [0, 100].
// 1 <= Node.val <= 100
// Node.val is unique for each node.
// There are no repeated edges and no self-loops in the graph.
// The Graph is connected and all nodes can be visited starting from the given node.

import (
	"fmt"
)

//Definition for a Node.
type Node struct {
	Val       int
	Neighbors []*Node
}

const (
	right = 0
	left  = 1
)

func cloneGraph(node *Node) *Node {
	if node == nil {
		return node
	} else if node.Neighbors == nil {
		return &(Node{Val: node.Val})
	} else if len(node.Neighbors) == 1 {
		curr := Node{Val: node.Val}
		next := Node{Val: node.Neighbors[right].Val}
		curr.Neighbors = []*Node{&next}
		next.Neighbors = []*Node{&curr}
		return &curr
	}

	nodeMap := make(map[int]*Node)
}

// [[2,3,4],[1,3,4],[1,2,4],[1,2,3]]

func main() {
	node1 := Node{Val: 1}
	node2 := Node{Val: 2}
	node3 := Node{Val: 3}
	node4 := Node{Val: 4}
	node1.Neighbors = []*Node{&node2, &node4}
	node2.Neighbors = []*Node{&node1, &node3}
	node3.Neighbors = []*Node{&node2, &node4}
	node4.Neighbors = []*Node{&node1, &node3}

	prevNodeVal := node4.Val

	for node := &node1; ; {
		fmt.Printf("%d ", node.Val)

		if node.Val == node4.Val {
			break

		}
		if node.Neighbors[0].Val == prevNodeVal {
			prevNodeVal = node.Val
			node = node.Neighbors[1]
		} else {
			prevNodeVal = node.Val
			node = node.Neighbors[0]
		}
	}

	fmt.Println(cloneGraph(&node1))
}
