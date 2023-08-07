package journey_to_the_moon

/*
 * Journey to the Moon
 *
 * https://www.hackerrank.com/challenges/journey-to-the-moon/problem?isFullScreen=false
 */

func journeyToMoon(n int32, astronaut [][]int32) int64 {
    if n == 1 {
        return 0
    }
    
    pairCount := int64(0)
    sameMap := make(map[int32]map[int32]bool)
    queue := []int32 { 0 }
    groupIndex := int32(0)
    groupMapper := make([]int32, n)
    
    for i := int32(0); i < n; i++ {
        sameMap[i] = make(map[int32]bool)
    }
    
    for _, pair := range astronaut {
        id1, id2 := pair[0], pair[1]    
        sameMap[id1][id2], sameMap[id2][id1] = true, true
    }
    
    for {
        if len(queue) == 0 {
            if len(sameMap) == 0 {
                break
            }
            
            groupIndex++
            
            for id := range sameMap {
                queue = append(queue, id)
                break
            }
        }
        
        targetId := queue[0]
        queue = queue[1:]
        groupMapper[targetId] = groupIndex
        for pairId := range sameMap[targetId] {
            queue = append(queue, pairId)
        }
        
        delete(sameMap, targetId)
    }
    
    for i := 0; i < len(groupMapper); i++ {
        for j := i + 1; j <len(groupMapper); j++ {
            if groupMapper[i] != groupMapper[j] {
                pairCount++
            }
        }
    }

    return pairCount
}