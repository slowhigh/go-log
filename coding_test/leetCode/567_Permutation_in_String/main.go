package main

import (
	"fmt"
	"strings"
)

// Given two strings s1 and s2, return true if s2 contains a permutation of s1, or false otherwise.

// In other words, return true if one of s1's permutations is the substring of s2.

// Example 1:

// Input: s1 = "ab", s2 = "eidbaooo"
// Output: true
// Explanation: s2 contains one permutation of s1 ("ba").
// Example 2:

// Input: s1 = "ab", s2 = "eidboaoo"
// Output: false

// Constraints:

// 1 <= s1.length, s2.length <= 104
// s1 and s2 consist of lowercase English letters.

func checkInclusion(s1 string, s2 string) bool {
	_s1 := s1
	var char string
	lenS1 := len(s1)
	lenS2 := len(s2)

	for i := 0; i < lenS2; i++ {
		for j := 0; j < lenS1; j++ {
			if i >= lenS2 {
				break
			}

			char = string(s2[i])

			if !strings.Contains(_s1, char) {
				if j != 0 {
					_s1 = s1
					i -= 2
				}

				break
			}

			_s1 = strings.Replace(_s1, char, "_", 1)
			i++

			if j == lenS1-1 {
				return true
			}
		}
	}

	return false
}

// abc, bbbca => true (bca)
// ab, eidbaooo => true (ba) => time out
// adc, dcda => true (cda)

// hello, ooolleoooleh => false
// ab, eidboaoo => false

func main() {
	s1 := "ab"
	s2 := "eidbaooo"
	fmt.Println(checkInclusion(s1, s2))
}
