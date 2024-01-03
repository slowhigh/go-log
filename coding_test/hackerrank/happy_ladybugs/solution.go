package happy_ladybugs

import (
	"sort"
	"strings"
)

// Happy Ladybugs
// https://www.hackerrank.com/challenges/happy-ladybugs/problem?isFullScreen=false

func happyLadybugs(b string) string {
	bArr := strings.Split(b, "")

	if strings.Contains(b, "_") {
		sort.Strings(bArr)
	}

	for i := 0; i < len(bArr); i++ {
		if bArr[i] == "_" {
			break
		}
		if i != 0 && bArr[i-1] == bArr[i] {
			continue
		}
		if i != len(bArr) - 1 && bArr[i] == bArr[i+1] {
			continue
		}

		return "NO"
	}

	return "YES"
}
