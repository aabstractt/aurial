package handler

import (
    "github.com/df-mc/dragonfly/server/player"
    "github.com/go-gl/mathgl/mgl64"
)

type MoveHandler interface {
    // HandleMove handles the movement of a player. If the movement should be cancelled, return true.
    // The new position, yaw and pitch are passed.
    // cancelled is true if the already was cancelled by another handler.
    HandleMove(p *player.Player, newPos mgl64.Vec3, newYaw, newPitch float64, cancelled bool) bool
}
