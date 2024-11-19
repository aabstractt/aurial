package player

import (
	"github.com/aabstractt/aurial"
	"github.com/aabstractt/aurial/player/context"
	"github.com/df-mc/dragonfly/server/player"
)

var itemUseRegistry = aurial.NewRegistry[ItemUseHandler]()

type ItemUseHandler interface {
	// HandleItemUse handles the player using an item.
	HandleItemUse(p *player.Player, ctx *context.ItemUseContext)
}

func ItemUseRegistry() *aurial.Registry[ItemUseHandler] {
	return itemUseRegistry
}
