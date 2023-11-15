package sherlock_and_anagrams

import (
	"sort"
	"strings"
)

func sherlockAndAnagrams(s string) int32 {
	str := strings.Split(s, "")
	sort.Strings(str)
	s = strings.Join(str, "")

	caseCount := make(map[string]int32)

	for i := 1; i <= len(s); i++ {
		for 
	}
}