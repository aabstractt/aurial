package player

import (
	"github.com/aabstractt/aurial"
	"github.com/aabstractt/aurial/player/context"
	"github.com/df-mc/dragonfly/server/player"
)

var hurtRegistry = aurial.NewRegistry[HurtHandler]()

type HurtHandler interface {
	// HandleHurt handles the player being hurt by any damage source.
	HandleHurt(p *player.Player, ctx *context.HurtContext)
}

func HurtRegistry() *aurial.Registry[HurtHandler] {
	return hurtRegistry
}
