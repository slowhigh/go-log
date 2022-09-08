package find_the_number_of_remaining_ones

func solution(n int) int {
    for x := 1; x < n - 1; x++ {
        if n % x == 1 {
            return x
        }
    }

	return n - 1
}