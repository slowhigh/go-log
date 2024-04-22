package larry_s_array

// Larry's Array
// https://www.hackerrank.com/challenges/larrys-array/problem?isFullScreen=false

func larrysArray(A []int32) string {
	count := 0
	for i := 0; i < len(A); i++ {
		for j := i + 1; j < len(A); j++ {
			if A[i] > A[j] {
				count++
			}
		}
	}

	if count%2 == 0 {
		return "YES"
	} else {
		return "NO"
	}
}
