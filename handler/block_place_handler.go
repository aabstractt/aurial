package handler

import (
    "github.com/df-mc/dragonfly/server/block/cube"
    "github.com/df-mc/dragonfly/server/player"
    "github.com/df-mc/dragonfly/server/world"
)

type BlockPlaceHandler interface {
    // BlockPlace is called when a block is placed by a player. The position of the block that is being
    BlockPlace(p *player.Player, pos cube.Pos, b world.Block) bool
}
