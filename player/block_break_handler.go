package player

import (
	"github.com/aabstractt/aurial"
	"github.com/aabstractt/aurial/player/context"
	"github.com/df-mc/dragonfly/server/player"
)

var blockBreakRegistry = aurial.NewRegistry[BlockBreakHandler]()

type BlockBreakHandler interface {
	// BlockBreak is called when a block is broken by a player. The context of the block break is passed
	// to the handler, which contains information about the block that was broken.
	BlockBreak(p *player.Player, ctx *context.BlockBreakContext)
}

func BlockBreakRegistry() *aurial.Registry[BlockBreakHandler] {
	return blockBreakRegistry
}
