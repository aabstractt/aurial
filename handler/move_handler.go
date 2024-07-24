package handler

import (
    "github.com/df-mc/dragonfly/server/event"
    "github.com/df-mc/dragonfly/server/player"
    "github.com/go-gl/mathgl/mgl64"
)

type MoveHandler interface {
    // HandleMove handles the movement of a player. ctx.Cancel() may be called to cancel the movement event.
    // The new position, yaw and pitch are passed.
    HandleMove(p *player.Player, ctx *event.Context, newPos mgl64.Vec3, newYaw, newPitch float64)
}
