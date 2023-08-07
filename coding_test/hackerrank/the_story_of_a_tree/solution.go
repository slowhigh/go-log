package the_story_of_a_tree

import "fmt"

/**
 * The Story of a Tree
 *
 * https://www.hackerrank.com/challenges/the-story-of-a-tree/problem?isFullScreen=false
 **/

func storyOfATree(n int32, edges [][]int32, k int32, guesses [][]int32) string {
	caseCountArr := make([]int32, n+1)
	edgeArr := make([][]int32, n+1)
	var visitedNodeArr []bool
	var result string

	for _, edge := range edges {
		n1, n2 := edge[0], edge[1]

		if edgeArr[n1] == nil {
			edgeArr[n1] = make([]int32, 0)
		}
		if edgeArr[n2] == nil {
			edgeArr[n2] = make([]int32, 0)
		}

		edgeArr[n1] = append(edgeArr[n1], n2)
		edgeArr[n2] = append(edgeArr[n2], n1)
	}

	queue := make([]int32, 0)
	for _, guess := range guesses {
		visitedNodeArr = make([]bool, n+1)
		parents, child := guess[0], guess[1]
		caseCountArr[parents]++
		visitedNodeArr[parents] = true

		for _, node := range edgeArr[parents] {
			if node != child {
				queue = append(queue, node)
			}
		}

		for len(queue) != 0 {
			from := queue[0]
			queue = queue[1:]
			caseCountArr[from]++
			visitedNodeArr[from] = true

			for _, to := range edgeArr[from] {
				if !visitedNodeArr[to] {
					queue = append(queue, to)
				}
			}
		}
	}

	rootCanCount := int32(0)
	for i := int32(1); i < n+1; i++ {
		if k <= caseCountArr[i] {
			rootCanCount++
		}
	}

	if rootCanCount == n {
		result = "1/1"
	} else if rootCanCount == 0 {
		result = "0/1"
	} else {
		for i := int32(2); i <= rootCanCount && i <= n; {
			if rootCanCount%i == 0 && n%i == 0 {
				rootCanCount /= i
				n /= i
			} else {
				i++
			}
		}

		result = fmt.Sprintf("%d/%d", rootCanCount, n)
	}

	return result
}
