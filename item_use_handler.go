package aurial

import (
	"github.com/aabstractt/aurial/context"
	"github.com/df-mc/dragonfly/server/player"
)

type ItemUseHandler interface {
	// HandleItemUse handles the player using an item.
	HandleItemUse(p *player.Player, ctx *context.ItemUseContext)
}
