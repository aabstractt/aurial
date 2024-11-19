package context

import "github.com/df-mc/dragonfly/server/world"

type ChangeWorldContext struct {
	before, after *world.World
}

// NewChangeWorldContext creates a new ChangeWorldContext with the world the player was in before the change
// and the world the player is in after the change.
func NewChangeWorldContext(before, after *world.World) *ChangeWorldContext {
	return &ChangeWorldContext{before, after}
}

// Before returns the world the player was in before the change.
func (ctx *ChangeWorldContext) Before() *world.World {
	return ctx.before
}

// After returns the world the player is in after the change.
func (ctx *ChangeWorldContext) After() *world.World {
	return ctx.after
}
