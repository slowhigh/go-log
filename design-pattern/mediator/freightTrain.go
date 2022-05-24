// Concrete component
package mediator

import "fmt"

type freightTrain struct {
	mediator mediator
}

func (g *freightTrain) arrive() bool {
	if !g.mediator.canArrive(g) {
		fmt.Println("FreightTrain : Arrival blocked, waiting")
		return false
	}

	fmt.Println("FreightTrain: Arrived")
	return true
}

func (g *freightTrain) depart() {
	fmt.Println("FreightTrain: Leaving")
	g.mediator.notifyAboutDeparture()
}

func (g *freightTrain) permitArrival() {
	fmt.Println("FreightTrain: Arrival permitted, arriving")
	g.arrive()
}