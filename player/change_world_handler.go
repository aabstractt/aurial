package player

import (
	"github.com/aabstractt/aurial"
	"github.com/aabstractt/aurial/player/context"
	"github.com/df-mc/dragonfly/server/player"
)

var changeWorldRegistry = aurial.NewRegistry[ChangeWorldHandler]()

type ChangeWorldHandler interface {
	// HandleChangeWorld handles when the player is added to a new world. before may be nil.
	HandleChangeWorld(p *player.Player, ctx *context.ChangeWorldContext)
}

func ChangeWorldRegistry() *aurial.Registry[ChangeWorldHandler] {
	return changeWorldRegistry
}
