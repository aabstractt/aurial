package handler

import (
    "github.com/df-mc/dragonfly/server/player"
    "github.com/go-gl/mathgl/mgl64"
)

type TeleportHandler interface {
    // HandleTeleport handles the teleportation of a player. If the teleportation should be cancelled, return true.
    // cancelled is true if the already was cancelled by another handler.
    HandleTeleport(p *player.Player, pos mgl64.Vec3, cancelled bool) bool
}
