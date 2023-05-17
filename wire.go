//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/thegodev/google-wire-demo/greeter"
	"github.com/thegodev/google-wire-demo/message"
)

// initializeApp creates an app with wire to inject all of the dependencies
func initializeApp(phrase string) (app, error) {
	wire.Build(provideApp, greeter.ProvideGreeter, message.ProvideMessage)
	return app{}, nil
}
