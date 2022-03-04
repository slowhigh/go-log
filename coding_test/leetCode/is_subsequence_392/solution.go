package is_subsequence_392

import (
	"strings"
)

func isSubsequence(s string, t string) bool {
	sArr, tArr := strings.Split(s, ""), strings.Split(t, "")
	sArrLen, tArrLen := len(sArr), len(tArr)

	if sArrLen == 0 {
		return true
	}

	for tIndex, sIndex := 0, 0; tIndex < tArrLen; tIndex++ {
		if tArr[tIndex] != sArr[sIndex] {
			continue
		}

		if sIndex+1 != sArrLen {
			sIndex++
			continue
		}

		return true
	}

	return false
}
