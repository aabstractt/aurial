package handler

import "github.com/df-mc/dragonfly/server/world"

type DeathHandler interface {
    // HandleDeath handles the player dying to a particular damage cause.
    HandleDeath(src world.DamageSource, keepInv *bool)
}
