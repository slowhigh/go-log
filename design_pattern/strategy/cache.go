// Context
package strategy

import "time"

type cache struct {
	storage      map[string]data
	evictionAlgo evictionAlgo
	capacity     int
	maxCapacity  int
}

func initCache(e evictionAlgo, maxCapacity int) *cache {
	storage := make(map[string]data)
	return &cache{
		storage:      storage,
		evictionAlgo: e,
		capacity:     0,
		maxCapacity:  maxCapacity,
	}
}

func (c *cache) setEvictionAlgo(e evictionAlgo) {
	c.evictionAlgo = e
}

func (c *cache) add(key string, value data) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}

	time.Sleep(1000)
	value.registeredTime = time.Now().UnixNano()
	c.capacity++
	c.storage[key] = value
}

func (c *cache) get(key string) {
	delete(c.storage, key)
}

func (c *cache) evict() {
	c.evictionAlgo.evict(c)
	c.capacity--
}

func (c *cache) contains(key string) bool {
	_, ok := c.storage[key]

	return ok
}
