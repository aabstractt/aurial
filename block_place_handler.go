package aurial

import (
	"github.com/aabstractt/aurial/context"
	"github.com/df-mc/dragonfly/server/player"
)

type BlockPlaceHandler interface {
	// HandleBlockPlace is called when a block is placed by a player.
	HandleBlockPlace(p *player.Player, ctx *context.BlockPlaceContext)
}
