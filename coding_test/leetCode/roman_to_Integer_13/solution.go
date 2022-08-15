package roman_to_integer_13

import "strings"

func romanToInt(s string) int {
	sum := 0

	sArr := strings.Split(s, "")
	before := ""

	for i := 0; i < len(sArr); i++ {
		c := sArr[i]
		n := before
		before = ""

		if n == "I" {
			if c == "V" {
				sum += 3
				continue
			} else if c == "X" {
				sum += 8
				continue
			}
		} else if n == "X" {
			if c == "L" {
				sum += 30
				continue
			} else if c == "C" {
				sum += 80
				continue
			}
		} else if n == "C" {
			if c == "D" {
				sum += 300
				continue
			} else if c == "M" {
				sum += 800
				continue
			}
		}

		if c == "M" {
			sum += 1000
		} else if c == "D" {
			sum += 500
		} else if c == "C" {
			sum += 100
		} else if c == "L" {
			sum += 50
		} else if c == "X" {
			sum += 10
		} else if c == "V" {
			sum += 5
		} else {
			sum += 1
		}

		before = c
	}

	return sum
}
