package contact

import "regexp"

func solution(signal string) string {
	if match, _ := regexp.MatchString("(100+1+|01)+", signal); match {
		return "YES"
	} else {
		return "NO"
	}
}
