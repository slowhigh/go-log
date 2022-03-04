package clone_graph_133

import (
	"testing"
)

func Test_cloneGraph(t *testing.T) {
	node1 := Node{Val: 1}
	node2 := Node{Val: 2}
	node3 := Node{Val: 3}
	node4 := Node{Val: 4}
	node1.Neighbors = []*Node{&node2, &node4}
	node2.Neighbors = []*Node{&node1, &node3}
	node3.Neighbors = []*Node{&node2, &node4}
	node4.Neighbors = []*Node{&node1, &node3}

	clonedGraph := cloneGraph(&node1)

	if clonedGraph.Val != node1.Val || clonedGraph.Neighbors[0].Val != node2.Val || clonedGraph.Neighbors[1].Val != node4.Val {
		t.Error("Wrong result!")
	}

	clonedGraph = clonedGraph.Neighbors[0]

	if clonedGraph.Val != node2.Val || clonedGraph.Neighbors[0].Val != node1.Val || clonedGraph.Neighbors[1].Val != node3.Val {
		t.Error("Wrong result!")
	}

	clonedGraph = clonedGraph.Neighbors[1]

	if clonedGraph.Val != node3.Val || clonedGraph.Neighbors[0].Val != node2.Val || clonedGraph.Neighbors[1].Val != node4.Val {
		t.Error("Wrong result!")
	}

	clonedGraph = clonedGraph.Neighbors[1]

	if clonedGraph.Val != node4.Val || clonedGraph.Neighbors[0].Val != node1.Val || clonedGraph.Neighbors[1].Val != node3.Val {
		t.Error("Wrong result!")
	}
}
