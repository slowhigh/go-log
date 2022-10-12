package number_string_and_english_word

import (
	"strconv"
	"strings"
)

func solution2(s string) int {
	f := strings.NewReplacer(
		"zero", "0",
		"one", "1",
		"two", "2",
		"three", "3",
		"four", "4",
		"five", "5",
		"six", "6",
		"seven", "7",
		"eight", "8",
		"nine", "9",
	)
	i, _ := strconv.Atoi(f.Replace(s))
	return i
}

func solution(s string) int {
	if sv, err := strconv.Atoi(s); err == nil {
		return sv
	}
	numStringArr := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for i, numStr := range numStringArr {
		if strings.Contains(s, numStr) {
			s = strings.ReplaceAll(s, numStr, strconv.Itoa(i))
		}
	}

	num, _ := strconv.Atoi(s)

	return num
}
