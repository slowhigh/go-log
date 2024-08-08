package bearAndSteadyGene

// Bear and Steady Gene
// https://www.hackerrank.com/challenges/bear-and-steady-gene/problem?isFullScreen=false

func steadyGene(gene string) int32 {
	min := int32(len(gene))
	avg := len(gene) / 4
	geneMap := map[byte]int{'G': 0, 'A': 0, 'C': 0, 'T': 0}
	overMap := make(map[byte]int)
	isOver := false

	for i := 0; i < len(gene); i++ {
		geneMap[gene[i]]++

		if avg < geneMap[gene[i]] {
			overMap[gene[i]]++
			isOver = true
		}
	}

	if !isOver {
		return 0
	}

	left := 0
	for right := 0; right < len(gene); right++ {
		if _, ok := overMap[gene[right]]; ok {
			overMap[gene[right]]--
		}

		for isExistMap(overMap) {
			if min > int32(right-left)+1 {
				min = int32(right-left) + 1
			}

			if _, ok := overMap[gene[left]]; ok {
				overMap[gene[left]]++
			}
			left++
		}
	}

	return min
}

func isExistMap(m map[byte]int) bool {
	for _, v := range m {
		if v > 0 {
			return false
		}
	}

	return true
}
