package negative_positive_number_add

func solution(absolutes []int, signs []bool) int {
	sum := 0

	for i := 0; i < len(absolutes); i++ {
		if signs[i] {
			sum += absolutes[i]
		} else {
			sum -= absolutes[i]
		}
	}

	return sum
}
