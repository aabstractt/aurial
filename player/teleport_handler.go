package player

import (
	"github.com/aabstractt/aurial"
	"github.com/aabstractt/aurial/player/context"
	"github.com/df-mc/dragonfly/server/player"
)

var teleportRegistry = aurial.NewRegistry[TeleportHandler]()

type TeleportHandler interface {
	// HandleTeleport handles the teleportation of a player. If the teleportation should be cancelled, return true.
	// cancelled is true if the already was cancelled by another handler.
	HandleTeleport(p *player.Player, ctx *context.TeleportContext)
}

// TeleportRegistry returns the registry of teleport handlers.
func TeleportRegistry() *aurial.Registry[TeleportHandler] {
	return teleportRegistry
}
