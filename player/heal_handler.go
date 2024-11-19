package player

import (
	"github.com/aabstractt/aurial"
	"github.com/aabstractt/aurial/player/context"
	"github.com/df-mc/dragonfly/server/player"
)

var healRegistry = aurial.NewRegistry[HealHandler]()

type HealHandler interface {
	// HandleHeal handles the player being healed by a healing source.
	HandleHeal(p *player.Player, ctx *context.HealContext)
}

func HealRegistry() *aurial.Registry[HealHandler] {
	return healRegistry
}
