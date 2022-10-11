package camouflage

func solution(clothes [][]string) int {
	closet := make(map[string]int)
	caseCount := 1

	for _, c := range clothes {
		closet[c[1]] += 1
	}

	for _, cou := range closet {
		caseCount *= cou + 1
	}

	return caseCount - 1
}
