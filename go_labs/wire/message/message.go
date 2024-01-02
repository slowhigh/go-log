package message

type Message string

func NewMessage() Message {
	return Message("Hi there!")
}

