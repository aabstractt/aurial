package handler

import (
    "github.com/df-mc/dragonfly/server/block/cube"
    "github.com/df-mc/dragonfly/server/item"
    "github.com/df-mc/dragonfly/server/player"
)

type BlockBreakHandler interface {
    // BlockBreak is called when a block is broken by a player. The position of the block that is being
    // cancelled is true if the already was cancelled by another handler.
    BlockBreak(p *player.Player, pos cube.Pos, drops *[]item.Stack, xp *int, cancelled bool) bool
}
