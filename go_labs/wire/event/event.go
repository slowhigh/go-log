package event

import (
	"fmt"
	"github/Slowhigh/go-study/go_labs/wire/greeter"
)

type Event struct {
	Greeter greeter.Greeter
}

func NewEvent(g greeter.Greeter) Event {
	return Event{Greeter: g}
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}