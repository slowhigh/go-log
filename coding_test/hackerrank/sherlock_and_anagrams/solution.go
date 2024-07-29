package sherlockAndAnagrams

import (
	"sort"
	"strings"
)

// Sherlock and Anagrams

// https://www.hackerrank.com/challenges/sherlock-and-anagrams/problem?isFullScreen=false

func sherlockAndAnagrams(s string) int32 {
	totalCount := int32(0)
	mapper := make(map[string]int32)

	for i := 1; i < len(s); i++ {
		for j := 0; j <= len(s)-i; j++ {
			sub := strings.Split(s[j:j+i], "")
			sort.Strings(sub)
			str := strings.Join(sub, "")
			if count, ok := mapper[str]; ok {
				mapper[str]++
				totalCount += count
			} else {
				mapper[str]++
			}
		}
	}

	return totalCount
}
