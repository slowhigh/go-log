package picking_numbers

// Picking Numbers
// https://www.hackerrank.com/challenges/picking-numbers/problem

func pickingNumbers(a []int32) int32 {
	numMap := make(map[int32]int32)
	topLen := int32(0)

	for _, num := range a {
		numMap[num]++
	}

	for k, v := range numMap {
		len := v + numMap[k-1]

		if topLen < len {
			topLen = len
		}
	}

	return topLen
}
