package handler

import (
    "github.com/df-mc/dragonfly/server/player"
    "github.com/df-mc/dragonfly/server/world"
)

type AttackEntityHandler interface {
    // HandleAttackEntity is called when a player attacks an entity. The entity that is being attacked is
    // passed, and the function may return true to cancel the attack.
    // cancelled is true if the already was cancelled by another handler.
    HandleAttackEntity(p *player.Player, e world.Entity, force, height *float64, critical *bool, cancelled bool) bool
}
