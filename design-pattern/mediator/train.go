// Component
package mediator

type train interface {
	arrive() bool
	depart()
	permitArrival()
}