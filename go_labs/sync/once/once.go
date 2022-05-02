package once

import (
	"sync"
)

type Counter struct {
	runCount int
	initCount int
	mutex sync.Mutex
	once sync.Once
}

func (c *Counter) Increment() {
	c.once.Do(func() {
		c.initCount += 1
	})

	c.mutex.Lock()
	c.runCount++
	c.mutex.Unlock()
}