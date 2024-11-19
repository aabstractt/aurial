package context

import "github.com/go-gl/mathgl/mgl64"

type TeleportContext struct {
	to mgl64.Vec3

	CancellableContext
}

// NewTeleportContext creates a new TeleportContext with the position to which the player is teleporting.
func NewTeleportContext(to mgl64.Vec3) *TeleportContext {
	return &TeleportContext{to: to}
}

// To returns the position to which the player is teleporting.
func (ctx *TeleportContext) To() mgl64.Vec3 {
	return ctx.to
}

// SetTo sets the position to which the player is teleporting.
func (ctx *TeleportContext) SetTo(to mgl64.Vec3) {
	ctx.to = to
}
