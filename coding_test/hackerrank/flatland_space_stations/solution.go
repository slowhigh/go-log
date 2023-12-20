package flatland_space_stations

import (
	"sort"
)

// Flatland Space Stations
// https://www.hackerrank.com/challenges/flatland-space-stations/problem?isFullScreen=false

func flatlandSpaceStations(n int32, c []int32) int32 {
	distArr := make([]int32, n)
	maxDist := int32(0)
	sort.Slice(c, func(i, j int) bool { return c[i] < c[j] })

	for i := 0; i < len(c); i++ {
		if c[i] > 0 {
			leftPin := int32(0)
			if i > 0 {
				leftPin = c[i-1] + 1
			}
			for l, d := c[i]-1, int32(1); l >= leftPin; l, d = l-1, d+1 {
				if distArr[l] == 0 || d < distArr[l] {
					distArr[l] = d
				}
			}
		}
		if c[i] < n-1 {
			rightPin := int32(n - 1)
			if i < len(c) - 1 {
				rightPin = c[i+1] - 1
			}
			for r, d := c[i]+1, int32(1); r <= rightPin; r, d = r+1, d+1 {
				if distArr[r] == 0 || d < distArr[r] {
					distArr[r] = d
				}
			}
		}
	}

	for _, dist := range distArr {
		if maxDist < dist {
			maxDist = dist
		}
	}

	return maxDist
}
