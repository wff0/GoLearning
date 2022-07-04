package wire

import (
	"errors"
	"fmt"
	"time"
)

type Message string

func NewMessage(phrase string) Message {
	return Message(phrase)
}

type Greeter struct {
	Message Message // <- adding a Message field
	Grumpy  bool
}

func (g Greeter) Greet() Message {
	if g.Grumpy {
		return "Go away!"
	}
	return g.Message
}

func NewGreeter(m Message) Greeter {
	var grumpy bool
	if time.Now().Unix()%2 == 0 {
		grumpy = true
	}
	return Greeter{Message: m, Grumpy: grumpy}
}

type Event struct {
	Greeter Greeter // <- adding a Greeter field
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

func NewEvent(g Greeter) (Event, error) {
	if g.Grumpy {
		return Event{}, errors.New("could not create event: event greeter is grumpy")
	}
	return Event{Greeter: g}, nil
}

func NewEventNumber() int {
	return 1
}
