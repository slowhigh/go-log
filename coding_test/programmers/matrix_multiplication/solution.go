package matrix_multiplication

func solution(arr1 [][]int, arr2 [][]int) [][]int {
	if arr1 == nil || arr2 == nil {
		return nil
	}

	x, y, z := len(arr1), len(arr2[0]), len(arr2)
	multiArr := make([][]int, x)

	for i := 0; i < x; i++ {
		multiArr[i] = make([]int, y)
		for j := 0; j < y; j++ {
			for k := 0; k < z; k++ {
				multiArr[i][j] += arr1[i][k] * arr2[k][j]
			}
		}
	}

	return multiArr
}
