package absolute_permutation

// Absolute Permutation
// https://www.hackerrank.com/challenges/absolute-permutation/problem?isFullScreen=false

func absolutePermutation(n int32, k int32) []int32 {
	p := make([]int32, n)
	isAdd := true

	if k == 0 {
		for i := int32(1); i <= n; i++ {
			p[i-1] = i
		}
	} else if n%(2*k) == 0 {
		for i := int32(1); i <= n; i++ {
			if isAdd {
				p[i-1] = i + k
			} else {
				p[i-1] = i - k
			}
			if i%k == 0 {
				isAdd = !isAdd
			}
		}
	} else {
		p = []int32{-1}
	}

	return p
}
