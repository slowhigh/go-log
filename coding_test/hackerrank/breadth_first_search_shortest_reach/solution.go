package breadth_first_search_shortest_reach

/**
 * Breadth First Search: Shortest Reach
 *
 * https://www.hackerrank.com/challenges/bfsshortreach/problem?isFullScreen=false
 **/

const DIST_UNIT = 6

func bfs(n int32, m int32, edges [][]int32, s int32) []int32 {
	distArr := make([]int32, n+1)
	edgeMap := make(map[int32][][]int32)

	for i := int32(1); i <= n; i++ {
		if i == s {
			distArr[i] = 0
		} else {
			distArr[i] = -1
		}

		edgeMap[i] = make([][]int32, 0)
	}

	for _, edge := range edges {
		src := edge[0]
		des := edge[1]

		edgeMap[src] = append(edgeMap[src], []int32{des, DIST_UNIT})
		edgeMap[des] = append(edgeMap[des], []int32{src, DIST_UNIT})
	}

	queue := make([]int32, 0)
	queue = append(queue, s)

	for len(queue) != 0 {
		src := queue[0]   // peek
		queue = queue[1:] // dequeue

		for _, edge := range edgeMap[src] {
			des := edge[0]
			dist := distArr[src] + edge[1]

			if distArr[des] == -1 || dist < distArr[des] {
				distArr[des] = dist
				queue = append(queue, des)
			}
		}
	}

	return append(distArr[1:s], distArr[s+1:]...)
}
