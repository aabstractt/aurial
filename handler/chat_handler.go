package handler

import (
    "github.com/df-mc/dragonfly/server/player"
)

type ChatHandler interface {
    // HandleChat handles a message sent in the chat by a player.
    // If the message should be cancelled, return true.
    // cancelled is true if the already was cancelled by another handler.
    HandleChat(p *player.Player, message *string, cancelled bool) bool
}
