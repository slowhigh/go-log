package personality_type_test

import "strings"

var (
	typeArr  = []string{"R", "T", "C", "F", "J", "M", "A", "N"}
	scoreArr = [][]int{{3, 0}, {2, 0}, {1, 0}, {0, 0}, {0, 1}, {0, 2}, {0, 3}}
)

func solution(survey []string, choices []int) string {
	typeScoreMap := make(map[string]int, 0)
	resultType := ""

	for _, t := range typeArr {
		typeScoreMap[t] = 0
	}

	for i := 0; i < len(survey); i++ {
		typeSplitArr := strings.Split(survey[i], "")

		typeScoreMap[typeSplitArr[0]] += scoreArr[choices[i]-1][0]
		typeScoreMap[typeSplitArr[1]] += scoreArr[choices[i]-1][1]
	}

	for i := 0; i < len(typeArr); i += 2 {
		if typeScoreMap[typeArr[i]] >= typeScoreMap[typeArr[i+1]] {
			resultType += typeArr[i]
		} else {
			resultType += typeArr[i+1]
		}
	}

	return resultType
}
