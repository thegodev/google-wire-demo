package message

// Message is what greeters will use to greet guests
type Message string

// ProvideMessage creates a default Message
func ProvideMessage(phrase string) Message {
	return Message(phrase)
}
