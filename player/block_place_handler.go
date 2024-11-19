package player

import (
	"github.com/aabstractt/aurial"
	"github.com/aabstractt/aurial/player/context"
	"github.com/df-mc/dragonfly/server/player"
)

var blockPlaceRegistry = aurial.NewRegistry[BlockPlaceHandler]()

type BlockPlaceHandler interface {
	// HandleBlockPlace is called when a block is placed by a player.
	HandleBlockPlace(p *player.Player, ctx *context.BlockPlaceContext)
}

func BlockPlaceRegistry() *aurial.Registry[BlockPlaceHandler] {
	return blockPlaceRegistry
}
