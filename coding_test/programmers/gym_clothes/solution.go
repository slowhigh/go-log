package gym_clothes

func solution(n int, lost []int, reserve []int) int {
	lostMap := make(map[int]struct{})
	for i := 0; i < len(lost); i++ {
		lostMap[lost[i]] = struct{}{}
	}

	// 여벌 있는 학생이 도난 당했을 때
	reserveMap := make(map[int]struct{})
	for i := 0; i < len(reserve); i++ {
		if _, exists := lostMap[reserve[i]]; exists {
			delete(lostMap, reserve[i])
		} else {
			reserveMap[reserve[i]] = struct{}{}
		}
	}

	lostCount := len(lostMap)

	// 여벌 있는 학생이 여벌을 빌려줄 때
	for k := range lostMap {
		if _, exists := reserveMap[k-1]; exists {
			delete(reserveMap, k-1)
			lostCount--
			continue
		}

		if _, exists := reserveMap[k+1]; exists {
			delete(reserveMap, k+1)
			lostCount--
		}
	}

	return n - lostCount
}
