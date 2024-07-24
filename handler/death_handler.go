package handler

import (
    "github.com/df-mc/dragonfly/server/player"
    "github.com/df-mc/dragonfly/server/world"
)

type DeathHandler interface {
    // HandleDeath handles the player dying to a particular damage cause.
    HandleDeath(p *player.Player, src world.DamageSource, keepInv *bool)
}
