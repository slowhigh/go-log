package the_grid_search

import "strings"

func gridSearch(G []string, P []string) string {
	fullG := strings.Join(G, "")
	fullGLen := len(fullG)
	pIndexArr := make([][]int, len(P))

	for i, p := range P {
		subFullG := fullG
		pIndexArr[i] = make([]int, 0)
		pin := 0

		for {
			findIndex := strings.Index(subFullG, p)
			if findIndex == -1 {
				break
			}

			pIndexArr[i] = append(pIndexArr[i], pin + findIndex)

			if fullGLen <= pin + findIndex+1 

			subFullG = subFullG[findIndex+1:]
		}
	}

	return ""
}
