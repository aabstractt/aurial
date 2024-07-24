package handler

import (
    "github.com/df-mc/dragonfly/server/block/cube"
    "github.com/df-mc/dragonfly/server/player"
    "github.com/df-mc/dragonfly/server/world"
)

type BlockPlaceHandler interface {
    // BlockPlace is called when a block is placed by a player. The position of the block that is being
    // cancelled is true if the already was cancelled by another handler.
    BlockPlace(p *player.Player, pos cube.Pos, b world.Block, cancelled bool) bool
}
