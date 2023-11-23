package three_d_surface_area

// 3D Surface Area
// https://www.hackerrank.com/challenges/3d-surface-area/problem?isFullScreen=false

func surfaceArea(A [][]int32) int32 {
	cost := int32(0)

	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			hight := A[i][j]
			cost += 2 // floor + roof

			//front
			if i == 0 {
				cost += hight
			} else if hight > A[i-1][j] {
				cost += hight - A[i-1][j]
			}

			//right
			if j == len(A[i])-1 {
				cost += hight
			} else if hight > A[i][j+1] {
				cost += hight - A[i][j+1]
			}

			//back
			if i == len(A)-1 {
				cost += hight
			} else if hight > A[i+1][j] {
				cost += hight - A[i+1][j]
			}

			//left
			if j == 0 {
				cost += hight
			} else if hight > A[i][j-1] {
				cost += hight - A[i][j-1]
			}
		}
	}

	return cost
}
