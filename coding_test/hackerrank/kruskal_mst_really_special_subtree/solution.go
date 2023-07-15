package kruskal_mst_really_special_subtree

import (
	"sort"
)

/*
 * Kruskal (MST): Really Special Subtree
 *
 * https://www.hackerrank.com/challenges/kruskalmstrsub/problem?isFullScreen=true
 */

type edge struct {
	from   int32
	to     int32
	weight int32
}

func kruskals(gNodes int32, gFrom []int32, gTo []int32, gWeight []int32) int32 {
	edgeCount := len(gWeight)
	edgeArr := make([]edge, edgeCount)
	cycleCheckArr := make([]int32, gNodes)
	weightSum := int32(0)

	for i := int32(0); i < gNodes; i++ {
		cycleCheckArr[i] = i
	}

	for i := 0; i < edgeCount; i++ {
		edgeArr[i] = edge{from: gFrom[i], to: gTo[i], weight: gWeight[i]}
	}

	sort.Slice(edgeArr, func(i, j int) bool { return edgeArr[i].weight < edgeArr[j].weight })

	for _, edge := range edgeArr {
		if cycleCheckArr[edge.from-1] == cycleCheckArr[edge.to-1] {
			continue
		}

		min, max := cycleCheckArr[edge.from-1], cycleCheckArr[edge.to-1]
		if max < min {
			min, max = max , min
		}

		for i := int32(0); i < gNodes; i++ {
			if cycleCheckArr[i] == max {
				cycleCheckArr[i] = min
			}
		}
		
		weightSum += edge.weight
	}

	return weightSum
}
