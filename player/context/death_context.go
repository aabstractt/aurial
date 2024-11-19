package context

import (
	"github.com/df-mc/dragonfly/server/block"
	"github.com/df-mc/dragonfly/server/entity"
	"github.com/df-mc/dragonfly/server/entity/effect"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

type DeathContext struct {
	src     world.DamageSource // The source of the damage that killed the player.
	keepInv bool               // Whether the player should keep their inventory after dying.

	messageCondition func(p *player.Player) bool
	message          string

	killer *player.Player
}

// NewDeathContext creates a new DeathContext with the source of the damage that killed the player.
func NewDeathContext(src world.DamageSource, keepInv bool) *DeathContext {
	return &DeathContext{src: src, keepInv: keepInv}
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

// Message returns the message that will be sent to the player when they die.
func (ctx *DeathContext) Message() string {
	return ctx.message
}

// SetMessage sets the message that will be sent to the player when they die.
func (ctx *DeathContext) SetMessage(message string) {
	ctx.message = message
}

// Killer returns the player that killed the player.
func (ctx *DeathContext) Killer() *player.Player {
	return ctx.killer
}

func (ctx *DeathContext) SetKiller(killer *player.Player) {
	ctx.killer = killer
}

// DamageName returns the name of the damage source that killed the player.
func DamageName(src world.DamageSource) string {
	if _, ok := src.(entity.AttackDamageSource); ok {
		return "attack"
	} else if _, ok = src.(entity.VoidDamageSource); ok {
		return "void"
	} else if _, ok = src.(entity.SuffocationDamageSource); ok {
		return "suffocation"
	} else if _, ok = src.(entity.DrowningDamageSource); ok {
		return "drowning"
	} else if _, ok = src.(entity.FallDamageSource); ok {
		return "fall"
	} else if _, ok = src.(entity.GlideDamageSource); ok {
		return "glide"
	} else if _, ok = src.(entity.LightningDamageSource); ok {
		return "lightning"
	} else if _, ok = src.(entity.ProjectileDamageSource); ok {
		return "projectile"
	} else if _, ok = src.(entity.ExplosionDamageSource); ok {
		return "explosion"
	} else if _, ok = src.(block.LavaDamageSource); ok {
		return "lava"
	} else if _, ok = src.(block.FireDamageSource); ok {
		return "fire"
	} else if srcBlock, ok := src.(block.DamageSource); ok {
		if bName, _ := srcBlock.Block.EncodeBlock(); bName != "" {
			return bName + "_block"
		}
	} else if _, ok = src.(effect.InstantDamageSource); ok {
		return "instant_damage"
	} else if _, ok = src.(effect.PoisonDamageSource); ok {
		return "poison"
	} else if _, ok = src.(effect.WitherDamageSource); ok {
		return "wither"
	} else if _, ok = src.(player.StarvationDamageSource); ok {
		return "hunger"
	}

	return "unknown"
}
