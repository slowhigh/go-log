package make_jadencase_string

import "strings"

func solution(s string) string {
	result := ""
	s = strings.ToLower(s)

	for _, str := range strings.Split(s, " ") {
		if str == "" {
			result += " "
			continue
		}

		if result != "" {
			result += " "
		}

		str = strings.Replace(str, string(str[0]), strings.ToUpper(string(str[0])), 1)
		result += str
	}

	return result
}
