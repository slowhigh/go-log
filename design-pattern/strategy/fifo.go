// Concrete strategy
package strategy

import "log"

type fifo struct{}

func (l *fifo) evict(c *cache) {
	evictedDataKey := ""

	for key, data := range c.storage {
		if evictedDataKey == "" {
			evictedDataKey = key
			continue
		}

		if data.registeredTime < c.storage[evictedDataKey].registeredTime {
			evictedDataKey = key
		}
	}

	delete(c.storage, evictedDataKey)
	log.Printf("Evict '%s' data.", evictedDataKey)
}
