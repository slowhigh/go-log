package roads_and_libraries

/**
 * Roads and Libraries
 *
 * https://www.hackerrank.com/challenges/torque-and-development/problem?isFullScreen=true
 **/

func roadsAndLibraries(n int32, c_lib int32, c_road int32, cities [][]int32) int64 {
	if len(cities) == 0 || c_lib <= c_road {
		return int64(c_lib) * int64(n)
	}

	minLib := int32(0)
	roadMap := getRoadMap(cities)
	minLib += n - int32(len(roadMap))
	minLib += getGroupCount(roadMap)

	return (int64(c_lib) * int64(minLib)) + (int64(c_road) * int64(n-minLib))
}

func getRoadMap(cities [][]int32) map[int32]map[int32]bool {
	roadMap := make(map[int32]map[int32]bool)

	for i := 0; i < len(cities); i++ {
		u, v := cities[i][0], cities[i][1]

		if len(roadMap[u]) == 0 {
			roadMap[u] = make(map[int32]bool, 0)
		}
		roadMap[u][v] = true

		if len(roadMap[v]) == 0 {
			roadMap[v] = make(map[int32]bool, 0)
		}
		roadMap[v][u] = true
	}

	return roadMap
}

func getGroupCount(roadMap map[int32]map[int32]bool) int32 {
	groupCount := int32(0)

	queue := []int32{}
	for {
		if len(queue) == 0 {
			if len(roadMap) == 0 {
				break
			}

			for k, _ := range roadMap {
				queue = append(queue, k)
				break
			}

			groupCount++
			continue
		}

		pivot := queue[0]

		if _, ok := roadMap[pivot]; ok {
			for k, _ := range roadMap[pivot] {
				delete(roadMap[k], pivot)
				queue = append(queue, k)
			}
		}

		delete(roadMap, pivot)
		queue = queue[1:]
	}

	return groupCount
}
