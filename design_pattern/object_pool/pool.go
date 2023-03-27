package object_pool

import (
	"fmt"
	"sync"
)

type pool struct {
	idle     []iPoolObject
	active   []iPoolObject
	capacity int
	mulock   *sync.Mutex
}

func InitPool(poolObject []iPoolObject) (*pool, error) {
	if len(poolObject) == 0 {
		return nil, fmt.Errorf("Cannot create a pool of 0 length")
	}

	pool := &pool{
		idle:     poolObject,
		active:   make([]iPoolObject, 0),
		capacity: len(poolObject),
		mulock:   new(sync.Mutex),
	}

	return pool, nil
}

func (p *pool) Loan() (iPoolObject, error) {
	p.mulock.Lock()
	defer p.mulock.Unlock()

	if len(p.idle) == 0 {
		return nil, fmt.Errorf("No pool object free. Please request after sometime")
	}

	obj := p.idle[0]
	p.idle = p.idle[1:]
	p.active = append(p.active, obj)
	fmt.Printf("Loan Pool Object with ID: %s\n", obj.getID())

	return obj, nil
}

func (p *pool) Receive(target iPoolObject) error {
	p.mulock.Lock()
	defer p.mulock.Unlock()

	err := p.remove(target)
	if err != nil {
		return err
	}

	p.idle = append(p.idle, target)

	fmt.Printf("Return Pool Object with ID: %s\n", target.getID())

	return nil
}

func (p *pool) remove(target iPoolObject) error {
	curActiveLength := len(p.active)

	for i, obj := range p.active {
		if obj.getID() == target.getID() {
			p.active[curActiveLength-1], p.active[i] = p.active[curActiveLength-1], p.active[i]
			p.active = p.active[:curActiveLength-1]
			return nil
		}
	}

	return fmt.Errorf("Target pool object doesn't belong to the pool")
}
