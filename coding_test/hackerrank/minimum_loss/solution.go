package minimum_loss

import (
	"math"
	"sort"
)

// Minimum Loss
// https://www.hackerrank.com/challenges/minimum-loss/problem?isFullScreen=false

type pack struct {
	Price int64
	Index int
}

func minimumLoss(price []int64) int32 {
	var (
		loss = int64(math.MaxInt64)
		arr  = make([]pack, 0)
	)

	for i, p := range price {
		arr = append(arr, pack{Price: p, Index: i})
	}

	sort.Slice(arr, func(i, j int) bool { return arr[i].Price > arr[j].Price })

	for i := 0; i < len(arr)-1; i++ {
		if arr[i].Index < arr[i+1].Index && arr[i].Price-arr[i+1].Price < loss {
			loss = arr[i].Price - arr[i+1].Price
		}
	}

	return int32(loss)
}
