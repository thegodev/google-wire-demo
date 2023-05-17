// The greeter binary simulates an event with greeters greeting guests.
package main

import (
	"fmt"
	"os"

	"github.com/camerontrew/DI-Test/greeter"
)

// newApp creates a new app - initialising all the dependencies
func provideApp(g greeter.Greeter) (app, error) {
	return app{Greeter: g}, nil
}

// here is a struct for your app with all the dependencies
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
