package commonChild

import (
	"log/slog"
	"strings"
)

// Common Child
// https://www.hackerrank.com/challenges/common-child/problem?isFullScreen=false

func commonChild(s1 string, s2 string) int32 {
	max := int32(0)
	slog.Info("=== new case ===", "s1", s1, "s2", s2)
	s1Arr, s2Arr := strings.Split(s1, ""), strings.Split(s2, "")

	newS1Arr, newS2Arr := make([]string, len(s1)), make([]string, len(s2))
	for i := 0; i < len(s1); i++ {
		if strings.Contains(s2, s1Arr[i]) {
			newS1Arr[i] = s1Arr[i]
		}

		if strings.Contains(s1, s2Arr[i]) {
			newS2Arr[i] = s2Arr[i]
		}
	}

	s1, s2 = strings.Join(newS1Arr, ""), strings.Join(newS2Arr, "")
	if len(s1) != 0 && len(s2) != 0 {
		slog.Info("new str", "s1", s1, "s2", s2)
		count, x, y, xPivot, yPivot := int32(0), 0, 0, 0, 0
		for {
			slog.Info("for", "count", count, "xPivot", xPivot, "yPivot", yPivot, "x", x, "y", y)
			if s1[x] == s2[y] {
				slog.Info("equal", "x", x, "y", y)
				count++
				x++
				y++
				yPivot = y
			} else {
				y++
			}

			if len(s2) <= y {
				x++
				y = yPivot
			}

			if len(s1) <= x || len(s2) <= yPivot {
				xPivot++
				x = xPivot
				yPivot = 0
				y = 0

				if max < count {
					max = count
				}

				count = 0
				slog.Info("update", "max", max)
			}

			if len(s1) <= xPivot {
				break
			}
		}
	}

	return max
}
