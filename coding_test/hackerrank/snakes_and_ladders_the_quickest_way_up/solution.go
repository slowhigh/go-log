package snakes_and_ladders_the_quickest_way_up

func quickestWayUp(ladders [][]int32, snakes [][]int32) int32 {
	edgeMap := make(map[int32][]int32)
	distArr := make([]int32, 101)
	nodeGroupArr := make([]int32, 101)
	queue := make([]int32, 0)

	for _, edge := range append(ladders, snakes...) {
		src, des := edge[0], edge[1]

		if _, ok := edgeMap[src]; !ok {
			edgeMap[src] = make([]int32, 0)
		}

		edgeMap[src] = append(edgeMap[src], des)
	}

	for i := int32(0); i < 101; i++ {
		distArr[i] = -1
		nodeGroupArr[i] = i
	}

	queue = append(queue, 1)
	distArr[1] = 0

	for len(queue) != 0 {
		src := queue[0]
		queue = queue[1:]

		for _, des := range edgeMap[src] {
			if nodeGroupArr[src] == nodeGroupArr[des] {
				continue
			}

			if distArr[des] == -1 || distArr[src] < distArr[des] {
				distArr[des] = distArr[src]

				for i := 0; i < 101; i++ {
					if nodeGroupArr[des] < nodeGroupArr[src] {
						nodeGroupArr[src] = nodeGroupArr[des]
					} else {
						
					}
				}

				queue = append(queue, des)
			}
		}

		for i := int32(1); i < 7; i++ {
			if distArr[src+i] == -1 || distArr[src]+1 < distArr[src+i] {
				distArr[src+i] = distArr[src] + 1

				queue = append(queue, src+i)
			}
		}
	}

	return distArr[100]
}
