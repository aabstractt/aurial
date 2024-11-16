package aurial

import (
	"github.com/aabstractt/aurial/context"
	"github.com/df-mc/dragonfly/server/player"
)

type HealHandler interface {
	// HandleHeal handles the player being healed by a healing source.
	HandleHeal(p *player.Player, ctx *context.HealContext)
}
