package context

import (
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

type DeathContext struct {
	src     world.DamageSource // The source of the damage that killed the player.
	keepInv bool               // Whether the player should keep their inventory after dying.

	messageCondition func(p *player.Player) bool
}

// Source returns the source of the damage that killed the player.
func (ctx *DeathContext) Source() world.DamageSource {
	return ctx.src
}

// KeepInventory returns whether the player should keep their inventory after dying.
func (ctx *DeathContext) KeepInventory() bool {
	return ctx.keepInv
}

// SetKeepInventory sets whether the player should keep their inventory after dying.
func (ctx *DeathContext) SetKeepInventory(keepInv bool) {
	ctx.keepInv = keepInv
}

// MessageCondition returns the condition that must be met for
// the death message to be sent to the player.
func (ctx *DeathContext) MessageCondition() func(p *player.Player) bool {
	return ctx.messageCondition
}

// SetMessageCondition sets the condition that must be met for
func (ctx *DeathContext) SetMessageCondition(condition func(p *player.Player) bool) {
	ctx.messageCondition = condition
}
