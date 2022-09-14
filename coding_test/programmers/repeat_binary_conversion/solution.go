package repeat_binary_conversion

import (
	"strconv"
	"strings"
)

func solution(s string) []int {
	repeatCount := 0
	var zeroCount int

	for repeatCount = 0; s != "1"; repeatCount++ {
		tempLen := len(s)
		s = strings.ReplaceAll(s, "0", "")
		zeroCount += tempLen - len(s)
		s = strconv.FormatInt(int64(len(s)), 2)
	}

	return []int{repeatCount, zeroCount}
}
