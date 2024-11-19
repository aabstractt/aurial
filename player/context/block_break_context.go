package context

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/item"
)

type BlockBreakContext struct {
	pos cube.Pos // The position of the block that is being broken.

	drops []item.Stack // The slice of item stacks that will be dropped when the block is broken.
	xp    int          // The amount of experience that will be dropped when the block is broken.

	CancellableContext
}

// NewBlockBreakContext creates a new BlockBreakContext with the position of the block that is being broken.
func NewBlockBreakContext(pos cube.Pos, drops []item.Stack, xp int) *BlockBreakContext {
	return &BlockBreakContext{
		pos: pos,

		drops: drops,
		xp:    xp,
	}
}

// Pos returns the position of the block that is being broken.
func (ctx *BlockBreakContext) Pos() cube.Pos {
	return ctx.pos
}

// Drops returns the slice of item stacks that will be dropped when the block is broken.
func (ctx *BlockBreakContext) Drops() []item.Stack {
	return ctx.drops
}

// SetDrops sets the slice of item stacks that will be dropped when the block is broken.
func (ctx *BlockBreakContext) SetDrops(drops []item.Stack) {
	ctx.drops = drops
}

// XP returns the amount of experience that will be dropped when the block is broken.
func (ctx *BlockBreakContext) XP() int {
	return ctx.xp
}

// SetXP sets the amount of experience that will be dropped when the block is broken.
func (ctx *BlockBreakContext) SetXP(xp int) {
	ctx.xp = xp
}
