package knightlOnAChessboard

// KnightL on a Chessboard
// https://www.hackerrank.com/challenges/knightl-on-chessboard/problem?isFullScreen=false

type move struct {
	Row  int32
	Col  int32
	Cost int32
}

func knightlOnAChessboard(n int32) [][]int32 {
	var (
		cost    = make([][]int32, n)
		minMove = make([][]int32, n-1)
	)
	for i := range cost {
		cost[i] = make([]int32, n)
	}
	for i := range minMove {
		minMove[i] = make([]int32, n-1)
	}

	for i := int32(1); i < n; i++ {
		for j := int32(i); j < n; j++ {
			for i := range cost {
				cost[i] = make([]int32, n)
			}

			direct := [8][2]int32{
				{+i, +j}, {+i, -j}, {-i, +j}, {-i, -j},
				{+j, +i}, {+j, -i}, {-j, +i}, {-j, -i},
			}

			queue := make([]move, 0)
			queue = append(queue, move{Row: 0, Col: 0, Cost: 0})

			for len(queue) > 0 {
				m := queue[0]
				queue = queue[1:]
				cost[m.Row][m.Col] = m.Cost

				for _, d := range direct {
					if 0 <= m.Row+d[0] && 0 <= m.Col+d[1] &&
						m.Row+d[0] < n && m.Col+d[1] < n &&
						(cost[m.Row+d[0]][m.Col+d[1]] == 0 || m.Cost+1 < cost[m.Row+d[0]][m.Col+d[1]]) {
						cost[m.Row+d[0]][m.Col+d[1]] = m.Cost + 1
						queue = append(queue, move{Row: m.Row + d[0], Col: m.Col + d[1], Cost: m.Cost + 1})
					}
				}
			}

			if cost[n-1][n-1] == 0 {
				cost[n-1][n-1] = -1
			}

			minMove[i-1][j-1] = cost[n-1][n-1]
			if i != j {
				minMove[j-1][i-1] = cost[n-1][n-1]
			}
		}
	}

	return minMove
}
