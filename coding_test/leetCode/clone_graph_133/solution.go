package clone_graph_133

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	} else if node.Neighbors == nil {
		return &(Node{Val: node.Val})
	}

	nodeMap := make(map[int]*Node)
	newNodeMap := make(map[int]*Node)

	queue := make([]*Node, 0)

	for curNode := node; ; {
		if _, isExisted := nodeMap[curNode.Val]; !isExisted {
			nodeMap[curNode.Val] = curNode
			queue = append(queue, curNode.Neighbors...)
		}

		if len(queue) == 0 {
			break
		}

		curNode = queue[0]
		queue = queue[1:]
	}

	for key, value := range nodeMap {
		newNodeMap[key] = &Node{Val: key, Neighbors: make([]*Node, len(value.Neighbors))}
	}

	for key, value := range nodeMap {
		for index, neighbor := range value.Neighbors {
			newNodeMap[key].Neighbors[index] = newNodeMap[neighbor.Val]
		}
	}

	return newNodeMap[node.Val]
}
