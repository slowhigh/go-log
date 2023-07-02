package synchronous_shopping

import (
	"math"
	"strconv"
	"strings"
)

type queueItem struct {
	centerName int32
	fishSum    int32
}

func maxInt32(x, y int32) int32 {
	return int32(math.Max(float64(x), float64(y)))
}

func shop(n int32, k int32, centers []string, roads [][]int32) int32 {
	const inf = int32(math.MaxInt32)
	roadsLen := len(roads)
	distArr := make([][]int32, n)
	maxFishSum := (1 << k) - 1
	fishSumArr := make([]int32, n)
	roadMap := make(map[int32]map[int32]int32)

	for i := 0; i < len(centers); i++ {
		for _, center := range strings.Split(centers[i], " ")[1:] {
			fish, _ := strconv.Atoi(center)
			fishSumArr[i] |= 1 << (fish - 1)
		}
	}

	for i := 0; i < roadsLen; i++ {
		p1, p2, dist := roads[i][0], roads[i][1], roads[i][2]
		if _, ok := roadMap[p1-1]; !ok {
			roadMap[p1-1] = make(map[int32]int32)
		}
		if _, ok := roadMap[p2-1]; !ok {
			roadMap[p2-1] = make(map[int32]int32)
		}

		roadMap[p1-1][p2-1], roadMap[p2-1][p1-1] = dist, dist
	}

	for i := int32(0); i < n; i++ {
		distArr[i] = make([]int32, maxFishSum+1)

		for j := 0; j < maxFishSum+1; j++ {
			distArr[i][j] = inf
		}
	}

	queue := make([]queueItem, 0)
	distArr[0][fishSumArr[0]] = 0
	queue = append(queue, queueItem{0, fishSumArr[0]})
	for 0 < len(queue) {
		qItem := queue[0]
		queue = queue[1:]
		p1 := qItem.centerName
		curFishSum := qItem.fishSum
		curDistSum := distArr[p1][curFishSum]
		for p2, p2Dist := range roadMap[p1] {
			newFishSum := curFishSum | fishSumArr[p2]
			minDist := distArr[p2][newFishSum]

			if minDist <= curDistSum+p2Dist {
				continue
			}

			minDist = curDistSum + p2Dist
			distArr[p2][newFishSum] = minDist
			queue = append(queue, queueItem{p2, newFishSum})
		}
	}

	twoCatMaxDist := inf

	for i := 0; i < maxFishSum+1; i++ {
		for j := 1; j < maxFishSum+1; j++ {
			if i|j == maxFishSum {
				dist := maxInt32(distArr[n-1][i], distArr[n-1][j])
				if dist < twoCatMaxDist {
					twoCatMaxDist = dist
				}
			}
		}
	}

	return twoCatMaxDist
}
