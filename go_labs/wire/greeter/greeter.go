package greeter

import "github/Slowhigh/go-study/go_labs/wire/message"

type Greeter struct {
	Message message.Message
}

func NewGreeter(m message.Message) Greeter {
	return Greeter{Message: m}
}

func (g Greeter) Greet() message.Message {
	return g.Message
}
