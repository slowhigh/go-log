package maximum_palindromes

import (
	"fmt"
)

// Maximum Palindromes

// https://www.hackerrank.com/challenges/maximum-palindromes/problem?isFullScreen=false

var str string
var strMap map[rune]int32

func initialize(s string) {
	str = s
	strMap = make(map[rune]int32)
	for _, r := range s {
		strMap[r] += 1
	}
}

func answerQuery(l int32, r int32) int32 {
	exMap := make(map[rune]int32)
	for _, r := range str[:l] {
		exMap[r] += 1
	}
	for _, r := range str[r:] {
		exMap[r] += 1
	}
	
	math.Fa

	for 

	fmt.Println(strMap)
	return 4
}