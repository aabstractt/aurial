package context

type ChatContext struct {
	message string

	CancellableContext
}

// NewChatContext creates a new ChatContext with the message that was sent in the chat.
func NewChatContext(message string) *ChatContext {
	return &ChatContext{message: message}
}

// Message returns the message that was sent in the chat.
func (ctx *ChatContext) Message() string {
	return ctx.message
}

// SetMessage sets the message that was sent in the chat.
func (ctx *ChatContext) SetMessage(message string) {
	ctx.message = message
}
