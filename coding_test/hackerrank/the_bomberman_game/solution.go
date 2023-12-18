package the_bomberman_game

import (
	"strings"
)

// The Bomberman Game
// https://www.hackerrank.com/challenges/bomber-man/problem?isFullScreen=false

func bomberMan(n int32, grid []string) []string {
	row, col := len(grid), len(grid[0])
	final := make([]string, row)

	if n == 1 {
		final = grid
	} else if n%2 == 0 {
		fill := strings.Repeat("O", col)
		for i := 0; i < row; i++ {
			final[i] = fill
		}
	} else {
		reverse := bomb(grid, row, col)

		if n%4 == 3 {
			final = reverse
		} else { // n%4 == 1
			final = bomb(reverse, row, col)
		}
	}

	return final
}

func bomb(grid []string, row, col int) []string {
	final := make([]string, row)
	reverse := make([][]string, row)
	reverse[0] = make([]string, col)

	for i := 0; i < row; i++ {
		if i < row-1 {
			reverse[i+1] = make([]string, col)
		}

		for j := 0; j < col; j++ {
			if string(grid[i][j]) == "." {
				if reverse[i][j] == "." {
					continue
				}

				reverse[i][j] = "O"
			} else {
				reverse[i][j] = "."

				// up
				if i > 0 {
					reverse[i-1][j] = "."
				}
				// down
				if i < row-1 {
					reverse[i+1][j] = "."
				}
				// left
				if j > 0 {
					reverse[i][j-1] = "."
				}
				// right
				if j < col-1 {
					reverse[i][j+1] = "."
				}
			}
		}

		for i := 0; i < row; i++ {
			final[i] = strings.Join(reverse[i], "")
		}
	}

	return final
}
