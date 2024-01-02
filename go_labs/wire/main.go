package main

import (
	"github/Slowhigh/go-study/go_labs/wire/event"
	"github/Slowhigh/go-study/go_labs/wire/greeter"
	"github/Slowhigh/go-study/go_labs/wire/message"
)

func main() {
	message := message.NewMessage()
	greeter := greeter.NewGreeter(message)
	event := event.NewEvent(greeter)

	event.Start()
}
