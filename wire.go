//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/thegodev/google-wire-demo/greeter"
	"github.com/thegodev/google-wire-demo/message"
)

// InitializeEvent creates an Event. It will error if the Event is staffed with
// a grumpy greeter.
func initializeApp(phrase string) (app, error) {
	wire.Build(provideApp, greeter.ProvideGreeter, message.ProvideMessage)
	return app{}, nil
}
