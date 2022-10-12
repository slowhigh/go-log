package n_least_common_multiple

func solution(arr []int) int {
	lcm := 1
	var tempArr []int

	for i := 2; ; {
		breakCount := 0
		tempArr = make([]int, len(arr))
		divCount := 0

		for j := 0; j < len(arr); j++ {
			if arr[j] >= i {
				breakCount++
			}

			if arr[j]%i == 0 {
				divCount++
				tempArr[j] = arr[j] / i
			} else {
				tempArr[j] = arr[j]
			}
		}

		if breakCount < 2 {
			break
		}

		if divCount > 1 {
			arr = tempArr
			lcm *= i
		} else {
			i++
		}
	}

	for _, v := range arr {
		lcm *= v
	}

	return lcm
}
