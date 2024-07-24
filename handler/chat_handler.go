package handler

import (
    "github.com/df-mc/dragonfly/server/event"
    "github.com/df-mc/dragonfly/server/player"
)

type ChatHandler interface {
    // HandleChat handles a message sent in the chat by a player. ctx.Cancel() may be called to cancel the
    // message being sent in chat.
    // The message may be changed by assigning to *message.
    HandleChat(p *player.Player, ctx *event.Context, message *string)
}
