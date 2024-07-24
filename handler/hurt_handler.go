package handler

import (
    "github.com/df-mc/dragonfly/server/player"
    "github.com/df-mc/dragonfly/server/world"
    "time"
)

type HurtHandler interface {
    // HandleHurt handles the player being hurt by any damage source. ctx.Cancel() may be called to cancel the
    // damage being dealt to the player.
    // The damage dealt to the player may be changed by assigning to *damage.
    // cancelled is true if the already was cancelled by another handler.
    HandleHurt(p *player.Player, damage *float64, attackImmunity *time.Duration, src world.DamageSource, cancelled bool) bool
}
