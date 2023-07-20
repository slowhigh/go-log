package remove_digit_from_number_to_maximize_result_2259

import (
	"strings"
)

func removeDigit(number string, digit byte) string {
	removeDigitIndex := strings.LastIndex(number, string(digit))
	numberLen := len(number)
	for i := 1; i < numberLen; i++ {
		if number[i-1] == digit && digit < number[i] {
			removeDigitIndex = i-1
			break
		}
	}

	return number[:removeDigitIndex] + number[removeDigitIndex+1:]
}
