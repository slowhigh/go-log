package highestValuePalindrome

import (
	"fmt"
	"math"
	"strings"
)

// Highest Value Palindrome
// https://www.hackerrank.com/challenges/richie-rich/problem?isFullScreen=false
// 0 < n < 10^5
// 0 <= k <= 10^5
func highestValuePalindrome(s string, n int32, k int32) string {
	sArr := strings.Split(s, "")
	check := make([]bool, n)
	var center int32
	if math.Mod(float64(n), float64(2)) == 0 {
		center = n/2 - 1
	} else {
		center = n / 2
	}

	fmt.Println("[1]")
	fmt.Printf("s = %s\n", s)
	fmt.Printf("center = %d\n", center)

	diff := 0
	for i := int32(0); i <= center; i++ {
		if sArr[i] != sArr[n-i-1] {
			diff++
		}
	}

	fmt.Println("[2]")
	fmt.Printf("diff = %d\n", diff)

	if k < int32(diff) {
		return "-1"
	} else if diff == 0 && k == 0 {
		return s
	}

	for i := int32(0); i <= center; i++ {
		if sArr[i] == sArr[n-i-1] {
			continue
		}

		if sArr[i] > sArr[n-i-1] {
			sArr[n-i-1] = sArr[i]
			check[n-i-1] = true
		} else {
			sArr[i] = sArr[n-i-1]
			check[i] = true
		}

		k -= 1
	}

	fmt.Println("[3]")
	fmt.Printf("k = %d\n", k)
	fmt.Printf("s = %s\n", strings.Join(sArr, ""))

	for i := int32(0); i <= center; i++ {
		l, r := sArr[i], sArr[n-i-1]
		cost := int32(0)
		if l != "9" {
			l = "9"

			if !check[i] {
				cost++
			}
		}
		if r != "9" && n-i-1 != i {
			r = "9"

			if !check[n-i-1] {
				cost++
			}
		}

		if cost <= k {
			if n-i-1 == i {
				sArr[i] = l
			} else {
				sArr[i], sArr[n-i-1] = l, r
			}

			k -= cost
		}
	}

	fmt.Println("[4]")
	fmt.Printf("k = %d\n", k)
	fmt.Printf("s = %s\n", strings.Join(sArr, ""))
	fmt.Println()
	fmt.Println()

	return strings.Join(sArr, "")
}
