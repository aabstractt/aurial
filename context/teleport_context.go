package context

import "github.com/go-gl/mathgl/mgl64"

type TeleportContext struct {
	to mgl64.Vec3
}

// To returns the position to which the player is teleporting.
func (ctx *TeleportContext) To() mgl64.Vec3 {
	return ctx.to
}

// SetTo sets the position to which the player is teleporting.
func (ctx *TeleportContext) SetTo(to mgl64.Vec3) {
	ctx.to = to
}
