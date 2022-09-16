package carpet

import "math"

func solution(brown int, yellow int) []int {
	x := ((4 + brown) + int(math.Sqrt(math.Pow(float64(4+brown), 2)-float64(16*brown)-float64(16*yellow)))) / 4
	y := (brown / 2) - x + 2

	if x < y {
		x, y = y, x
	}

	return []int{x, y}
}
