package cavity_map

// Cavity Map
// https://www.hackerrank.com/challenges/cavity-map/problem?isFullScreen=false

func cavityMap(grid []string) []string {
	gridLen := len(grid)

	for i := 1; i < gridLen-1; i++ {
		for j := 1; j < gridLen-1; j++ {
			if grid[i][j] > grid[i-1][j] && grid[i][j] > grid[i][j+1] && grid[i][j] > grid[i+1][j] && grid[i][j] > grid[i][j-1] {
				grid[i] = grid[i][:j] + "X" + grid[i][j+1:]
			}
		}
	}

	return grid
}
