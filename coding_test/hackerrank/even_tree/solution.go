package even_tree

/*
 * Even Tree
 *
 * https://www.hackerrank.com/challenges/even-tree/problem
 */

func evenForest(t_nodes int32, t_edges int32, t_from []int32, t_to []int32) int32 {
	removedEdgeCount := int32(0)
	parentArr := make([]int32, t_nodes + 1)
	childCountArr := make([]int32, t_nodes + 1)

	for i := int32(0); i < t_edges; i++ {
		if t_from[i] < t_to[i] {
			parentArr[t_to[i]] = t_from[i]
		} else {
			parentArr[t_from[i]] = t_to[i]
		}
	}

	startNode := int32(2)
	targetNode := int32(2)
	for startNode <= t_nodes {
		childCountArr[targetNode]++
		targetNode = parentArr[targetNode]

		if targetNode == 1 {
			startNode++
			targetNode = startNode
		}
	}

	for i := int32(2); i <= t_nodes; i++ {
		if childCountArr[i] % 2 == 0 {
			removedEdgeCount++
		}
	}

	return removedEdgeCount
}
