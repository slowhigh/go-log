package k_th_number

import "sort"

func solution(array []int, commands [][]int) []int {
	result := []int{}

	for _, v := range commands {
		slice := make([]int, v[1]-v[0]+1)
		copy(slice, array[v[0]-1:v[1]])
		sort.Ints(slice)
		result = append(result, slice[v[2]-1])
	}

	return result
}
