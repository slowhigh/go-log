package n_least_common_multiple

func solution(arr []int) int {
	min := 0
	arrLen := len(arr)
	sum := make([]int, arrLen)
	count := make(map[int]int)

	for i := 0; ; i = (i + 1)%arrLen {
		sum[i] += arr[i]
		count[sum[i]]++

		if count[sum[i]] == arrLen {
			min = sum[i]
			break
		}
	}

	return min
}
