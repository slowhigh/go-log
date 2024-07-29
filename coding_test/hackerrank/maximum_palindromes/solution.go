package maximumPalindromes

import (
	"fmt"
	"math"
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
	caseCount := int32(1)
	duplicate := make([]int32, 0)
	totalTuple := int32(0)
	single := int32(0)

	fmt.Printf("str[:l-1] = %+v\n", str[:l-1])
	fmt.Printf("str[r:] = %+v\n", str[r:])


	for _, r := range str[:l-1] {
		exMap[r] += 1
	}
	for _, r := range str[r:] {
		exMap[r] += 1
	}

	fmt.Printf("strMap = %+v\n", strMap)
	fmt.Printf("exMap = %+v\n", exMap)

	for k, v := range strMap {
		count := v - exMap[k]
		tuple := count / 2
		totalTuple += tuple
		duplicate = append(duplicate, tuple)
		single += count % 2
	}

	fmt.Printf("totalTuple = %d\n", totalTuple)
	fmt.Printf("single = %d\n", single)
	fmt.Printf("duplicate = %+v\n", duplicate)
	fmt.Println()
	fmt.Println()

	if totalTuple == 0 {
		return 0
	}

	for i := totalTuple; 2 <= i; i-- {
		caseCount *= i
	}

	for _, d := range duplicate {
		for i := d; 2 <= i; i-- {
			caseCount /= i
		}
	}

	if (single != 0) {
		caseCount *= single
	}

	
	return caseCount % int32(math.Pow10(9) + 7)
}
