package find_the_prime_number_in_k_series

import (
	"strconv"
	"strings"
)

func solution(n int, k int) int {
	kNumStr := ""
	for n != 0 {
		q := n / k
		r := n - (q * k)
		kNumStr = strconv.Itoa(r) + kNumStr
		n = q
	}

	primeNumMap := make(map[string]interface{})
	count := 0
	for _, str := range strings.FieldsFunc(kNumStr, func(r rune) bool { return r == '0' }) {
		if str == "1" {
			continue
		}
		if _, ok := primeNumMap[str]; ok {
			count++
			continue
		}
		num, err := strconv.Atoi(str)
		if err != nil {
			continue
		}

		isPrimeNumber := true
		var max int
		max = num/2
		for j := 2; j <= max; j++ {
			if num%j == 0 {
				isPrimeNumber = false
				break
			}
			max = num/j
		}

		if isPrimeNumber {
			primeNumMap[str] = nil
			count++
		}
	}

	return count
}
