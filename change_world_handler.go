package aurial

import (
	"github.com/aabstractt/aurial/context"
	"github.com/df-mc/dragonfly/server/player"
)

type ChangeWorldHandler interface {
	// HandleChangeWorld handles when the player is added to a new world. before may be nil.
	HandleChangeWorld(p *player.Player, ctx *context.ChangeWorldContext)
}
