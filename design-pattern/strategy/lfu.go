// Concrete strategy
package strategy

import "log"

type lfu struct{}

func (l *lfu) evict(c *cache) {
	evictedDataKey := ""

	for key, data := range c.storage {
		if evictedDataKey == "" {
			evictedDataKey = key
			continue
		}

		if data.usedCount < c.storage[evictedDataKey].usedCount {
			evictedDataKey = key
		}
	}

	delete(c.storage, evictedDataKey)
	log.Printf("Evict '%s' data.", evictedDataKey)
}
