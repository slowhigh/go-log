package first_unique_character_in_a_string_387

import (
	"strings"
)

func firstUniqChar(s string) int {
	sArr := strings.Split(s, "")
	find := ""

	for i := 0; i < len(sArr); i++ {
		if strings.Contains(find, sArr[i]) {
			continue
		} else if i-1 >= 0 && strings.Contains(s[:i-1], sArr[i]) {
			find += sArr[i]
			continue
		} else if i+1 < len(sArr) && strings.Contains(s[i+1:], sArr[i]) {
			find += sArr[i]
			continue
		}

		return i
	}

	return -1
}
