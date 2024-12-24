package context

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/go-gl/mathgl/mgl64"
)

type MoveContext struct {
	to  mgl64.Vec3
	rot cube.Rotation

	CancellableContext
}

func NewMoveContext(to mgl64.Vec3, rot cube.Rotation) *MoveContext {
	return &MoveContext{
		to:  to,
		rot: rot,
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

// Rotation returns the rotation to which the player is moving.
func (ctx *MoveContext) Rotation() cube.Rotation {
	return ctx.rot
}

// SetRotation sets the rotation to which the player is moving.
func (ctx *MoveContext) SetRotation(rot cube.Rotation) {
	ctx.rot = rot
}
