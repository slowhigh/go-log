package main

// 점심시간에 도둑이 들어, 일부 학생이 체육복을 도난당했습니다. 다행히 여벌 체육복이 있는 학생이 이들에게 체육복을 빌려주려 합니다. 학생들의 번호는 체격 순으로 매겨져 있어, 바로 앞번호의 학생이나 바로 뒷번호의 학생에게만 체육복을 빌려줄 수 있습니다. 예를 들어, 4번 학생은 3번 학생이나 5번 학생에게만 체육복을 빌려줄 수 있습니다. 체육복이 없으면 수업을 들을 수 없기 때문에 체육복을 적절히 빌려 최대한 많은 학생이 체육수업을 들어야 합니다.

// 전체 학생의 수 n, 체육복을 도난당한 학생들의 번호가 담긴 배열 lost, 여벌의 체육복을 가져온 학생들의 번호가 담긴 배열 reserve가 매개변수로 주어질 때, 체육수업을 들을 수 있는 학생의 최댓값을 return 하도록 solution 함수를 작성해주세요.

// 제한사항
// 전체 학생의 수는 2명 이상 30명 이하입니다.
// 체육복을 도난당한 학생의 수는 1명 이상 n명 이하이고 중복되는 번호는 없습니다.
// 여벌의 체육복을 가져온 학생의 수는 1명 이상 n명 이하이고 중복되는 번호는 없습니다.
// 여벌 체육복이 있는 학생만 다른 학생에게 체육복을 빌려줄 수 있습니다.
// 여벌 체육복을 가져온 학생이 체육복을 도난당했을 수 있습니다. 이때 이 학생은 체육복을 하나만 도난당했다고 가정하며, 남은 체육복이 하나이기에 다른 학생에게는 체육복을 빌려줄 수 없습니다.
// 입출력 예
// n	lost	reserve		return
// 5	[2, 4]	[1, 3, 5]	5
// 5	[2, 4]	[3]			4
// 3	[3]		[1]			2
// 입출력 예 설명
// 예제 #1
// 1번 학생이 2번 학생에게 체육복을 빌려주고, 3번 학생이나 5번 학생이 4번 학생에게 체육복을 빌려주면 학생 5명이 체육수업을 들을 수 있습니다.

// 예제 #2
// 3번 학생이 2번 학생이나 4번 학생에게 체육복을 빌려주면 학생 4명이 체육수업을 들을 수 있습니다.

import (
	"fmt"
)

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

//6,[6,2,4],[1,5,3]
// n: 3, lost: [1, 2], reserve: [2, 3]
//	5, [2, 4], [3]

//6,[1,3,5],[2,4,6]

//n = 5, lost = [3, 5], reserve = [2, 4], return =5

func main() {
	n := 5
	lost := []int{1, 3}
	reserve := []int{2, 4}
	fmt.Println(solution(n, lost, reserve))
}
