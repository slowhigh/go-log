package not_exist_numbers_add

func solution(numbers []int) int {
	sum := 0 + 1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9

	for i := 0; i < len(numbers); i++ {
		sum -= numbers[i]
	}

	return sum
}
