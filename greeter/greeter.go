package greeter

import (
	"github.com/thegodev/google-wire-demo/message"
)

// ProvideGreeter initializes a Greeter
func ProvideGreeter(m message.Message) Greeter {
	return Greeter{Message: m}
}

// Greeter is the type for greeting guests
type Greeter struct {
	Message message.Message
}

// Greet produces a greeting for guests
func (g Greeter) Greet() message.Message {
	return g.Message
}
