package snakesAndLaddersTheQuickestWayUp

/**
 * Snakes and Ladders: The Quickest Way Up
 *
 * https://www.hackerrank.com/challenges/the-quickest-way-up/problem?isFullScreen=false
 **/

func quickestWayUp(ladders [][]int32, snakes [][]int32) int32 {
	edgeMap := make(map[int32]int32)
	distArr := make([]int32, 101)
	queue := make([]int32, 0)

	for _, edge := range append(ladders, snakes...) {
		edgeMap[edge[0]] = edge[1]
	}

	for i := int32(0); i < 101; i++ {
		distArr[i] = -1
	}

	queue = append(queue, 1)
	distArr[1] = 0

	for len(queue) != 0 {
		src := queue[0]
		queue = queue[1:]

		if des, ok := edgeMap[src]; ok && (distArr[des] == -1 || distArr[src] < distArr[des]) {
			distArr[des] = distArr[src]

			queue = append(queue, des)
		} else {
			for i := int32(1); i < 7 && src+i < 101; i++ {
				if distArr[src+i] == -1 || distArr[src]+1 < distArr[src+i] {
					distArr[src+i] = distArr[src] + 1

					queue = append(queue, src+i)
				}
			}
		}
	}

	return distArr[100]
}
