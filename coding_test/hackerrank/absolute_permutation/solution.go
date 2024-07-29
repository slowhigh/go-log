package absolutePermutation

import "math"

// Absolute Permutation
// https://www.hackerrank.com/challenges/absolute-permutation/problem?isFullScreen=false

func absolutePermutation(n int32, k int32) []int32 {
	p := make([]int32, n)

	if k == 0 {
		for i := int32(1); i <= n; i++ {
			p[i-1] = i
		}
	} else if n%2 != 0 || n/2 < k || math.Mod(float64(n)/float64(k), 2) != 0 {
		p = []int32{-1}
	} else {
		for i := int32(0); i < n; i++ {
			if (i/k)%2 == 0 {
				p[i] = (i + 1) + k
			} else {
				p[i] = (i + 1) - k
			}
		}
	}

	return p
}
