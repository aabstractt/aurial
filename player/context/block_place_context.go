package context

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/world"
)

type BlockPlaceContext struct {
	pos cube.Pos    // The position at which the block is being placed.
	b   world.Block // The block that is being placed.

	CancellableContext
}

// NewBlockPlaceContext creates a new BlockPlaceContext with the given position and block.
func NewBlockPlaceContext(pos cube.Pos, b world.Block) *BlockPlaceContext {
	return &BlockPlaceContext{
		pos: pos,
		b:   b,
	}
}

// Pos returns the position at which the block is being placed.
func (ctx *BlockPlaceContext) Pos() cube.Pos {
	return ctx.pos
}

// Block returns the block that is being placed.
func (ctx *BlockPlaceContext) Block() world.Block {
	return ctx.b
}
