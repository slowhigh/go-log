package special_subtree

import (
	"sort"
)

// Prim's (MST) : Special Subtree
//
// https://www.hackerrank.com/challenges/primsmstsub/problem?isFullScreen=true

type Node struct {
	no     int32
	weight int32
}

func prims(n int32, edges [][]int32, start int32) int32 {
	totalWeight := int32(0)
	edgeMap := make(map[int32]map[int32]int32)
	priorityQueue := make([]Node, 0)
	visitedNodeArr := make(map[int32]bool)
	visitedNodeCount := int32(0)

	for _, edge := range edges {
		node1, node2, weight := edge[0], edge[1], edge[2]

		if _, ok := edgeMap[node1]; !ok {
			edgeMap[node1] = make(map[int32]int32)
		}
		edgeMap[node1][node2] = weight

		if _, ok := edgeMap[node2]; !ok {
			edgeMap[node2] = make(map[int32]int32)
		}
		edgeMap[node2][node1] = weight
	}

	visitedNodeArr[1] = true
	for to, weight := range edgeMap[1] {
		priorityQueue = append(priorityQueue, Node{no: to, weight: weight})
	}

	for visitedNodeCount < n {
		var from Node

		sort.Slice(priorityQueue, func(i, j int) bool { return priorityQueue[i].weight < priorityQueue[j].weight })
		for i, node := range priorityQueue {
			if _, ok := visitedNodeArr[node.no]; !ok {
				from = Node{ no: node.no, weight: node.weight}
				priorityQueue = priorityQueue[i+1:]
				break
			}
		}

		visitedNodeArr[from.no] = true
		visitedNodeCount++
		totalWeight += from.weight

		for to, weight := range edgeMap[from.no] {
			if _, ok := visitedNodeArr[to]; !ok {
				priorityQueue = append(priorityQueue, Node{no: to, weight: weight})
			}
		}
	}

	return totalWeight
}
