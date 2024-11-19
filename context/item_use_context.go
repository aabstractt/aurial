package context

import "github.com/df-mc/dragonfly/server/item"

type ItemUseContext struct {
	handStack item.Stack // handStack is the item stack at his main hand.
	leftStack item.Stack // leftStack is the item stack at his left hand.

	CancellableContext
}

// NewItemUseContext creates a new ItemUseContext with the item stack that is being used.
func NewItemUseContext(handStack, leftStack item.Stack) *ItemUseContext {
	return &ItemUseContext{handStack: handStack, leftStack: leftStack}
}

// HandStack returns the item stack at the player's main hand.
func (ctx *ItemUseContext) HandStack() item.Stack {
	return ctx.handStack
}

// LeftStack returns the item stack at the player's left hand.
func (ctx *ItemUseContext) LeftStack() item.Stack {
	return ctx.leftStack
}
