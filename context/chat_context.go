package context

type ChatContext struct {
	message string

	CancellableContext
}

// Message returns the message that was sent in the chat.
func (ctx *ChatContext) Message() string {
	return ctx.message
}

// SetMessage sets the message that was sent in the chat.
func (ctx *ChatContext) SetMessage(message string) {
	ctx.message = message
}
