package context

import "github.com/df-mc/dragonfly/server/world"

type ChangeWorldContext struct {
	before, after *world.World
}

// Before returns the world the player was in before the change.
func (ctx *ChangeWorldContext) Before() *world.World {
	return ctx.before
}

// After returns the world the player is in after the change.
func (ctx *ChangeWorldContext) After() *world.World {
	return ctx.after
}
