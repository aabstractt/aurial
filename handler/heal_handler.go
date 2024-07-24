package handler

import (
    "github.com/df-mc/dragonfly/server/event"
    "github.com/df-mc/dragonfly/server/player"
    "github.com/df-mc/dragonfly/server/world"
)

type HealHandler interface {
    // HandleHeal handles the player being healed by a healing source. ctx.Cancel() may be called to cancel
    // the healing.
    // The health added may be changed by assigning to *health.
    HandleHeal(p *player.Player, ctx *event.Context, health *float64, src world.HealingSource)
}
