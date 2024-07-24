package handler

import (
    "github.com/df-mc/dragonfly/server/player"
    "github.com/df-mc/dragonfly/server/world"
)

type HealHandler interface {
    // HandleHeal handles the player being healed by a healing source. If the health should be cancelled, return true.
    // The health added may be changed by assigning to *health.
    // cancelled is true if the already was cancelled by another handler.
    HandleHeal(p *player.Player, health *float64, src world.HealingSource, cancelled bool) bool
}
