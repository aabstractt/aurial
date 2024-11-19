package player

import (
	"github.com/aabstractt/aurial"
	"github.com/aabstractt/aurial/player/context"
	"github.com/df-mc/dragonfly/server/player"
)

var deathRegistry = aurial.NewRegistry[DeathHandler]()

type DeathHandler interface {
	// HandleDeath handles the player dying to a particular damage cause.
	HandleDeath(p *player.Player, ctx *context.DeathContext)
}

// DeathRegistry returns the registry for death handlers.
func DeathRegistry() *aurial.Registry[DeathHandler] {
	return deathRegistry
}
