// Client code
package strategy

import (
	"testing"
)

func Test_Strategy(t *testing.T) {
	const maxCapacity = 2

	lfu := &lfu{}
	lru := &lru{}
	fifo := &fifo{}

	cache := initCache(lfu, maxCapacity)
	cache.add("a", data{waitingTime: 2, usedCount: 4})
	cache.add("b", data{waitingTime: 1, usedCount: 3})

	if !cache.contains("a") || !cache.contains("b") {
		t.Error("wrong add function")
	}

	cache.add("c", data{waitingTime: 3, usedCount: 2})

	if !cache.contains("a") || cache.contains("b") || !cache.contains("c") {
		t.Error("wrong lfu evict function")
	}

	cache.setEvictionAlgo(lru)
	cache.add("d", data{waitingTime: 4, usedCount: 5})

	if cache.contains("a") || !cache.contains("c") || !cache.contains("d") {
		t.Error("wrong lru evict function")
	}

	cache.setEvictionAlgo(fifo)
	cache.add("e", data{waitingTime: 5, usedCount: 1})

	if cache.contains("c") || !cache.contains("d") || !cache.contains("e") {
		t.Error("wrong fifo evict function")
	}
}
