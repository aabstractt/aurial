package handler

import (
    "github.com/df-mc/dragonfly/server/event"
    "github.com/df-mc/dragonfly/server/player"
    "github.com/go-gl/mathgl/mgl64"
)

type TeleportHandler interface {
    // HandleTeleport handles the teleportation of a player. ctx.Cancel() may be called to cancel it.
    HandleTeleport(p *player.Player, ctx *event.Context, pos mgl64.Vec3)
}
