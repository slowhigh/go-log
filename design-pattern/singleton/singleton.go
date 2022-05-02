package sigleton

import "sync"

type Singleton interface {
	AddCount(value int)
	GetCount() int
}

///
// Example 1 - sync
///
type singleton_1 struct {
	count int
}

var instance_1 *singleton_1

func GetInstance_1() Singleton {

	if instance_1 == nil {
		instance_1 = new(singleton_1)
	}

	return instance_1
}

func (s *singleton_1) AddCount(value int) {
	s.count += value
}

func (s *singleton_1) GetCount() int {
	return s.count
}

///
// Example 2 - async
///
var lock = &sync.Mutex{}

type singleton_2 struct {
	count int
}

var instance_2 *singleton_2

func GetInstance_2() Singleton {

	lock.Lock()
	defer lock.Unlock()

	if instance_2 == nil {
		instance_2 = new(singleton_2)
	}

	return instance_2
}

func (s *singleton_2) AddCount(value int) {
	s.count += value
}

func (s *singleton_2) GetCount() int {
	return s.count
}

///
// Example 3 - async
///

var once sync.Once

type singleton_3 struct {
	count int
}

var instance_3 *singleton_3

func GetInstance_3() Singleton {
	once.Do(func() {
		instance_3 = new(singleton_3)
	})

	return instance_3
}

func (s *singleton_3) AddCount(value int) {
	s.count += value
}

func (s *singleton_3) GetCount() int {
	return s.count
}
