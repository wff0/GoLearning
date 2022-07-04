//go:build wireinject
// +build wireinject

package wire

import "github.com/google/wire"

func InitializeEvent(phrase string) (Event, error) {
	wire.Build(NewEvent, NewGreeter, NewMessage, NewEventNumber)
	return Event{}, nil
}
