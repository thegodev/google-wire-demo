//go:build wireinject
// +build wireinject

package main

import (
	"github.com/camerontrew/DI-Test/greeter"
	"github.com/camerontrew/DI-Test/message"
	"github.com/google/wire"
)

// InitializeEvent creates an Event. It will error if the Event is staffed with
// a grumpy greeter.
func initializeApp(phrase string) (app, error) {
	wire.Build(provideApp, greeter.ProvideGreeter, message.ProvideMessage)
	return app{}, nil
}
