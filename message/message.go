package message

// Message is what greeters will use to greet guests
type Message string

// ProvideMessage initializes a Message
func ProvideMessage(phrase string) Message {
	return Message(phrase)
}
