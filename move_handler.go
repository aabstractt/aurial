package aurial

import (
	"github.com/aabstractt/aurial/context"
	"github.com/df-mc/dragonfly/server/player"
)

type MoveHandler interface {
	// HandleMove handles the movement of a player.
	HandleMove(p *player.Player, ctx *context.MoveContext)
}
