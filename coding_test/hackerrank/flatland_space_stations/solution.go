package flatland_space_stations

import (
	"sort"
)

// Flatland Space Stations
// https://www.hackerrank.com/challenges/flatland-space-stations/problem?isFullScreen=false

func flatlandSpaceStations(n int32, c []int32) int32 {
	distArr := make([]int32, n)
	maxDist := int32(0)
	sort.Slice(c, func(i, j int) bool { return i < j })

	for i := 0; i < len(c); i++ {
		if c[i] > 0 {
			// left
			for l, d := c[i]-1, int32(1); l > c[i-1]; l, d = l-1, d+1 {
				if d < distArr[l] {
					distArr[l] = d

					if maxDist < d {
						maxDist = d
					}
				}
			}
		}
		if c[i] < n-1 {
			// right
			for r, d := c[i]+1, int32(1); r < c[i+1]; r, d = r+1, d+1 {
				if d < distArr[r] {
					distArr[r] = d

					if maxDist < d {
						maxDist = d
					}
				}
			}
		}
	}

	return maxDist
}
