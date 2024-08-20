package contact

import (
	"fmt"
	"regexp"
)

func solution(signal string) string {
	match, err := regexp.MatchString(`^(100+1+|01)+$`, signal)
	if err != nil {
		fmt.Println("Error:", err)
		return "NO"
	}
	if match {
		return "YES"
	} else {
		return "NO"
	}
}
