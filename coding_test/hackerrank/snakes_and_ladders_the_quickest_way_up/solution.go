package snakes_and_ladders_the_quickest_way_up

func quickestWayUp(ladders [][]int32, snakes [][]int32) int32 {
	edgeMap := make(map[int32][]int32)
	distArr := make([]int32, 101)
	queue := make([]int32, 0)

	for _, ladder := range ladders {
		src := ladder[0]
		des := ladder[1]

		if _, ok := edgeMap[src]; !ok {
			edgeMap[src] = make([]int32, 0)
		}

		edgeMap[src] = append(edgeMap[src], des)
		queue = append(queue, src)
	}

	for _, snake := range snakes {
		src := snake[1]
		des := snake[0]

		if _, ok := edgeMap[src]; !ok {
			edgeMap[src] = make([]int32, 0)
		}

		edgeMap[src] = append(edgeMap[src], des)
		queue = append(queue, src)
	}

	for i := int32(1); i <= 100; i++ {
		distArr[i] = ((i - 1) / 6) + 1
	}

	for len(queue) != 0 {
		src := queue[0]
		queue = queue[1:]
		for _, des := range edgeMap[src] {
			newDist := distArr[src]

			if newDist < distArr[des] {
				distArr[des] = newDist

				if _, ok := edgeMap[des]; ok {
					queue = append(queue, des)
				}
			}

			for i := int32(1); des + i <= 100; i++ {
				newDist2 := newDist + ((i - 1) / 6) + 1

				if newDist2 < distArr[des + i] {
					distArr[des + i] = newDist2

					if _, ok := edgeMap[des + i]; ok {
						queue = append(queue, des + i)
					}
				}
			}
		}
	}

	return distArr[100]
}
