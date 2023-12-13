package the_bomberman_game

import (
	"fmt"
	"strings"
)

// The Bomberman Game
// https://www.hackerrank.com/challenges/bomber-man/problem?isFullScreen=false

func bomberMan(n int32, grid []string) []string {
	row, col := len(grid), len(grid[0])
	final := make([]string, row)

	fmt.Printf("row %d col %d\n", row, col)

	if n%2 == 0 {
		fill := strings.Repeat("O", col)
		for i := 0; i < col; i++ {
			final[0] = fill
		}
	} else if (n-1)%4 == 0 {
		final = grid
	} else {
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
					fmt.Printf("i+1=%d, j=%d\n", i+1, j)
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
	}

	return final
}
