package strange_counter

import (
	"fmt"
	"math"
)

// Strange Counter
// https://www.hackerrank.com/challenges/strange-code/problem?isFullScreen=false

func strangeCounter(t int64) int64 {
	// Write your code here

	fmt.Println(math.Log2(float64(t) / float64(3)))

	return 0
}
