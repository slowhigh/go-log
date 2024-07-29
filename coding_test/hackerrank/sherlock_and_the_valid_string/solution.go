package sherlockAndTheValidString

/* Sherlock and the Valid String
 *
 * https://www.hackerrank.com/challenges/sherlock-and-valid-string/problem?isFullScreen=false
 *
 */
func isValid(s string) string {
	caseMap := make(map[rune]int)
	countMap := make(map[int]int)

	for _, r := range s {
		caseMap[r] += 1
	}

	for _, c := range caseMap {
		countMap[c] += 1
	}

	if len(countMap) == 1 {
		return "YES"
	} else if len(countMap) > 2 {
		return "NO"
	} else {
		keys := make([]int, 0, 2)
		for k := range countMap {
			keys = append(keys, k)
		}

		if countMap[keys[0]] == 1 && (keys[0] == 1 || keys[0]+1 == keys[1] || keys[0]-1 == keys[1]) {
			return "YES"
		} else if countMap[keys[1]] == 1 && (keys[1] == 1 || keys[1]+1 == keys[0] || keys[1]-1 == keys[0]) {
			return "YES"
		} else {
			return "NO"
		}
	}
}
