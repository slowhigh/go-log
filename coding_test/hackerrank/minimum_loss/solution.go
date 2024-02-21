package minimum_loss

import "math"

// Minimum Loss
// https://www.hackerrank.com/challenges/minimum-loss/problem?isFullScreen=false

func minimumLoss(price []int64) int32 {
	var (
		loss = int32(math.MaxInt32)
	)

	for i := 0; i < len(price)-1; i++ {
		for j := i + 1; j < len(price); j++ {
			if price[i] >= price[j] && int32(price[i]-price[j]) < loss {
				loss = int32(price[i] - price[j])
			}
		}
	}

	return loss
}
