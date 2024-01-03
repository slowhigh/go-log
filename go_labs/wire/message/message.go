package message

import "github.com/google/wire"

type Message string

func NewMessage() Message {
	return Message("Hi there!")
}

var MessageSet = wire.NewSet(NewMessage)