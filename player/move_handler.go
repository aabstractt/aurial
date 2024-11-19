package player

import (
	"github.com/aabstractt/aurial"
	"github.com/aabstractt/aurial/player/context"
	"github.com/df-mc/dragonfly/server/player"
)

var moveRegistry = aurial.NewRegistry[MoveHandler]()

type MoveHandler interface {
	// HandleMove handles the movement of a player.
	HandleMove(p *player.Player, ctx *context.MoveContext)
}

func MoveRegistry() *aurial.Registry[MoveHandler] {
	return moveRegistry
}
