// Concrete strategy
package strategy

import "log"

type lru struct {}

func (l *lru) evict(c *cache) {
	evictedDataKey := ""
	
	for key, data := range c.storage {
		if evictedDataKey == "" {
			evictedDataKey = key
			continue
		}
		
		if data.waitingTime < c.storage[evictedDataKey].waitingTime {
			evictedDataKey = key
		}
	}
	
	delete(c.storage, evictedDataKey)
	log.Printf("Evict '%s' data.", evictedDataKey)
}