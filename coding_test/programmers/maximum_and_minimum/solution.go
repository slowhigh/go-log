package maximum_and_minimum

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func solution(s string) string {
	stringArr := strings.Split(s, " ")
	stringArrLen := len(stringArr)
	numArr := make([]int, stringArrLen)

	for i := 0; i < stringArrLen; i++ {
		number, err := strconv.Atoi(stringArr[i])
		if err != nil {
			break
		}

		numArr[i] = number
	}

	sort.Ints(numArr)

	return fmt.Sprintf("%d %d", numArr[0], numArr[stringArrLen-1])
}
