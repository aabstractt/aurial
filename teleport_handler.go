package aurial

import (
	"github.com/aabstractt/aurial/context"
	"github.com/df-mc/dragonfly/server/player"
)

type TeleportHandler interface {
	// HandleTeleport handles the teleportation of a player. If the teleportation should be cancelled, return true.
	// cancelled is true if the already was cancelled by another handler.
	HandleTeleport(p *player.Player, ctx *context.TeleportContext)
}
