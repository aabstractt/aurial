package context

import "github.com/go-gl/mathgl/mgl64"

type MoveContext struct {
	to mgl64.Vec3

	yaw   float64
	pitch float64

	CancellableContext
}

func NewMoveContext(to mgl64.Vec3, yaw, pitch float64) *MoveContext {
	return &MoveContext{
		to: to,

		yaw:   yaw,
		pitch: pitch,
	}
}

// To returns the position to which the player is moving.
func (ctx *MoveContext) To() mgl64.Vec3 {
	return ctx.to
}

// SetTo sets the position to which the player is moving.
func (ctx *MoveContext) SetTo(to mgl64.Vec3) {
	ctx.to = to
}

// Yaw returns the yaw to which the player is moving.
func (ctx *MoveContext) Yaw() float64 {
	return ctx.yaw
}

// SetYaw sets the yaw to which the player is moving.
func (ctx *MoveContext) SetYaw(yaw float64) {
	ctx.yaw = yaw
}

// Pitch returns the pitch to which the player is moving.
func (ctx *MoveContext) Pitch() float64 {
	return ctx.pitch
}

// SetPitch sets the pitch to which the player is moving.
func (ctx *MoveContext) SetPitch(pitch float64) {
	ctx.pitch = pitch
}
