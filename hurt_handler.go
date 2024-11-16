package aurial

import (
	"github.com/aabstractt/aurial/context"
	"github.com/df-mc/dragonfly/server/player"
)

type HurtHandler interface {
	// HandleHurt handles the player being hurt by any damage source.
	HandleHurt(p *player.Player, ctx *context.HurtContext)
}
