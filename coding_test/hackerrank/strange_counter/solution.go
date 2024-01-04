package strange_counter

import (
	"math"
)

// Strange Counter
// https://www.hackerrank.com/challenges/strange-code/problem?isFullScreen=false

func strangeCounter(t int64) int64 {
	var (
		sum         = int64(0)
		cycleEndNum = int64(0)
	)

	for i := float64(0); ; i++ {
		cycleEndNum += int64(3 * math.Pow(2, i))

		if t <= cycleEndNum {
			sum = cycleEndNum + 1
			break
		}
	}

	return sum - t
}
