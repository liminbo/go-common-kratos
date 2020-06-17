// +build wireinject
// The build tag makes sure the stub is not built in the final build.
package example_1

import "github.com/google/wire"

func initEvent(msg string) Event{
	wire.Build(newMessage, newGreeter,newEvent)
	return Event{}
}

