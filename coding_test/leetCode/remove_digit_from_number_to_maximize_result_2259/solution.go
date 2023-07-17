package remove_digit_from_number_to_maximize_result_2259

import (
	"strconv"
)

func removeDigit(number string, digit byte) string {
	maxNum := 0
	targetStr := string(digit)
	for i, r := range number {
		if string(r) == targetStr {
			number[:i]
		}
	}

	return strconv.Itoa(maxNum)
}
