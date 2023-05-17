package main

import (
	"fmt"
	"os"

	"github.com/thegodev/google-wire-demo/greeter"
)

// provideApp returns a new app with all of the dependencies
func provideApp(g greeter.Greeter) (app, error) {
	return app{Greeter: g}, nil
}

// app is the struct for your app with all the dependencies
type app struct {
	Greeter greeter.Greeter
}

// start the app with all your dependencies
func (a app) start() {
	msg := a.Greeter.Greet()
	fmt.Println(msg)
}

func main() {
	app, err := initializeApp("hi there!")
	if err != nil {
		fmt.Printf("failed to create app: %s\n", err)
		os.Exit(2)
	}
	app.start()
}
