package aurial

import (
	"github.com/aabstractt/aurial/context"
	"github.com/df-mc/dragonfly/server/player"
)

type DeathHandler interface {
	// HandleDeath handles the player dying to a particular damage cause.
	HandleDeath(p *player.Player, ctx *context.DeathContext)
}
