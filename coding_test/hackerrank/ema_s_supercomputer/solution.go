package emaSSupercomputer

import (
	"log/slog"
	"strings"
)

// Ema's Supercomputer
// https://www.hackerrank.com/challenges/two-pluses/problem?isFullScreen=false

const Good = "G"
const Bad = "B"

func twoPluses(grid []string) int32 {
	var result int32
	row, col := len(grid), len(grid[0])
	gridArr := make([][]string, row)

	for i := 0; i < row; i++ {
		gridArr[i] = strings.Split(grid[i], "")
	}

	max := row
	if col < max {
		max = col
	}
	max = (max - 1) / 2
	slog.Info("cal max", "row", row, "col", col, "max", max)

	tempArr := make([][]string, row)
	var count int
	var mul int
	for m := max; m >= 0; m-- {
		copy2D(tempArr, gridArr)
		count = 0
		mul = 1
		for size := m; size >= 0; size-- {
			for y := 0 + size; y < row-size; y++ {
				for x := 0 + size; x < col-size; x++ {
					slog.Info("check", "m", m, "size", size, "y", y, "x", x)
					if checkCell(row, col, tempArr, x, y, size) {
						count++
						mul *= 1 + (4 * size)
						slog.Info("==== check true", "count", count, "nul", mul)
					}

					if count >= 2 {
						goto BREAK
					}
				}
			}
		}

	BREAK:
		slog.Info("compare", "result", result, "mul", mul)
		if result < int32(mul) {
			result = int32(mul)
		}
		slog.Info("compare result", "result", result)
	}

	return result
}

func checkCell(row, col int, grid [][]string, x, y, size int) bool {
	slog.Info("checkCell param", "grid", grid)
	temp := make([][]string, row)
	copy2D(temp, grid)

	if temp[y][x] != Good {
		return false
	}

	for i := 1; i <= size; i++ {
		if x-i < 0 || temp[y][x-i] != Good {
			return false
		}
		temp[y][x-i] = Bad

		if y-i < 0 || temp[y-i][x] != Good {
			return false
		}
		temp[y-i][x] = Bad

		if col < x+i-1 || temp[y][x+i] != Good {
			return false
		}
		temp[y][x+i] = Bad

		if row < y+i-1 || temp[y+i][x] != Good {
			return false
		}
		temp[y+i][x] = Bad
	}

	copy2D(grid, temp)
	return true
}

func copy2D(dst, src [][]string) {
	for i, s := range src {
		dst[i] = make([]string, len(s))
		copy(dst[i], s)
	}
} 
