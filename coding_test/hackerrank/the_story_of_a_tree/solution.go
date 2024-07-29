package theStoryOfATree

import (
	"fmt"
)

// The Story of a Tree
//
// https://www.hackerrank.com/challenges/the-story-of-a-tree/problem?isFullScreen=false
func storyOfATree(n int32, edges [][]int32, k int32, guesses [][]int32) string {
	caseCountArr := make([]int32, n+1)
	edgeArr := make([][]int32, n+1)
	guessMap := make(map[int32]map[int32]bool, 0)
	queue := make([]int32, 0)
	rootCanCount := int32(0)
	var result string

	for _, edge := range edges {
		n1, n2 := edge[0], edge[1]

		if edgeArr[n1] == nil {
			edgeArr[n1] = make([]int32, 0)
		}
		if edgeArr[n2] == nil {
			edgeArr[n2] = make([]int32, 0)
		}

		edgeArr[n1], edgeArr[n2] = append(edgeArr[n1], n2), append(edgeArr[n2], n1)
	}

	for _, guess := range guesses {
		from, to := guess[1], guess[0]

		if _, ok := guessMap[from]; !ok {
			guessMap[from] = make(map[int32]bool)
		}

		guessMap[from][to] = true
	}

	visited := make([]bool, n+1)
	queue = append(queue, 1)

	for len(queue) > 0 {
		from := queue[0]
		queue = queue[1:]
		visited[from] = true

		for _, to := range edgeArr[from] {
			if visited[to] {
				continue
			}

			if _, ok := guessMap[to][from]; ok {
				caseCountArr[1]++
			}

			queue = append(queue, to)
		}
	}

	queue = append(queue, 1)
	visited = make([]bool, n+1)

	for len(queue) > 0 {
		from := queue[0]
		queue = queue[1:]
		visited[from] = true

		for _, to := range edgeArr[from] {
			if visited[to] {
				continue
			}

			caseCountArr[to] = caseCountArr[from]

			if _, ok := guessMap[from][to]; ok {
				caseCountArr[to] += 1
			}
			if _, ok := guessMap[to][from]; ok {
				caseCountArr[to] -= 1
			}

			queue = append(queue, to)
		}
	}

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
