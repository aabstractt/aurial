package player

import (
	"github.com/aabstractt/aurial"
	"github.com/aabstractt/aurial/player/context"
	"github.com/df-mc/dragonfly/server/player"
)

var chatRegistry = aurial.NewRegistry[ChatHandler]()

type ChatHandler interface {
	// HandleChat handles a message sent in the chat by a player.
	// If the message should be cancelled, return true.
	// cancelled is true if the already was cancelled by another handler.
	HandleChat(p *player.Player, ctx *context.ChatContext)
}

// ChatRegistry returns the registry for chat handlers.
func ChatRegistry() *aurial.Registry[ChatHandler] {
	return chatRegistry
}
