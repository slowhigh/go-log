// Concrete component
package mediator

import "fmt"

type passengerTrain struct {
	mediator mediator
}

func (g *passengerTrain) arrive() bool {
	if !g.mediator.canArrive(g) {
		fmt.Println("PassengerTrain : Arrival blocked, waiting")
		return false
	}

	fmt.Println("PassengerTrain: Arrived")
	return true
}

func (g *passengerTrain) depart() {
	fmt.Println("PassengerTrain: Leaving")
	g.mediator.notifyAboutDeparture()
}

func (g *passengerTrain) permitArrival() {
	fmt.Println("PassengerTrain: Arrival permitted, arriving")
	g.arrive()
}
