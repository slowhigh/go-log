package commonChild

// Common Child
// https://www.hackerrank.com/challenges/common-child/problem?isFullScreen=false
// LCS Algorithm
// Longest Common Subsequence
func commonChild(s1 string, s2 string) int32 {
	dp := make([][]int32, len(s1)+1)

	for x := 0; x < len(dp); x++ {
		dp[x] = make([]int32, len(s2)+1)

		for y := 0; y < len(s2)+1; y++ {
			if x == 0 || y == 0 {
				continue
			} else if s1[x-1] == s2[y-1] {
				dp[x][y] = dp[x-1][y-1] + 1
			} else {
				dp[x][y] = maxInt32(dp[x-1][y], dp[x][y-1])
			}
		}
	}

	return dp[len(s1)][len(s2)]
}

func maxInt32(a, b int32) int32 {
	if a > b {
		return a
	} else {
		return b
	}
}
