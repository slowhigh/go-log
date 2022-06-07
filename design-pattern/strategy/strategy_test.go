// Client code
package strategy

import "testing"

func Test_Strategy(t *testing.T) {
	const maxCapacity = 2

	lfu := &lfu{}
	lru := &lru{}
	fifo := &fifo{}

	cache := initCache(lfu, maxCapacity)
	cache.add("a", data{ waitingTime: 2, frequencyUse: 4})
	cache.add("b", data{ waitingTime: 1, frequencyUse: 3})
	
	if !cache.contains("a") || !cache.contains("b") {
		t.Error("wrong add function")
	}

	cache.add("c", data{ waitingTime: 3, frequencyUse: 2})

	

	cache.add("d", data{ waitingTime: 4, frequencyUse: 5})
	cache.add("e", data{ waitingTime: 5, frequencyUse: 1})
}