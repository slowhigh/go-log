package rotate_parentheses

import "strings"

func solution(s string) int {
	sArr := strings.Split(s, "")
	lenS := len(sArr)

	if lenS <= 1 {
		return 0
	}

	count := 0

	for i := 0; i < lenS; i++ {
		stack := make([]string, 0)

		for j := i; j < lenS+i; j++ {
			target := sArr[j%lenS]
			if target == "[" || target == "{" || target == "(" {
				stack = append(stack, target)
				continue
			}

			var flip string
			if target == "]" {
				flip = "["
			} else if target == "}" {
				flip = "{"
			} else {
				flip = "("
			}

			if len(stack) == 0 || stack[len(stack)-1] != flip {
				stack = append(stack, target)
				break
			}

			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			count++
		}
	}

	return count
}
