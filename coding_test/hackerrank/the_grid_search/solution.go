package the_grid_search


/*
 * The Grid Search
 *
 * https://www.hackerrank.com/challenges/the-grid-search/problem?isFullScreen=true
 */
func gridSearch(G []string, P []string) string {
	R, C := len(G), len(G[0])
	r, c := len(P), len(P[0])

	for i := 0; i <= R-r; i++ {
		for j := 0; j <= C-c; j++ {
			isMatch := true
			for k := 0; k < r; k++ {
				if G[i+k][j:j+c] != P[k] {
					isMatch = false
					break
				}
			}

			if isMatch {
				return "YES"
			}
		}
	}

	return "NO"
}
