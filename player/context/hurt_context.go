package context

import (
	"github.com/df-mc/dragonfly/server/entity"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"time"
)

type HurtContext struct {
	// Damage is the amount of damage that is dealt to the player. This value may be changed to change the
	// amount of damage dealt to the player.
	damage float64

	// Immune is whether the player is immune to the damage that is dealt to them. If the player is immune, the
	// damage will not be dealt to the player.
	immune bool

	// AttackImmunity is the duration for which the player is immune to attacks. This value may be changed to
	// change the duration of immunity.
	attackImmunity time.Duration

	// Source is the source of the damage that is dealt to the player.
	src world.DamageSource

	CancellableContext
}

// NewHurtContext creates a new HurtContext with the damage, attack immunity and source passed.
func NewHurtContext(damage float64, immune bool, attackImmunity time.Duration, src world.DamageSource) *HurtContext {
	return &HurtContext{
		damage:         damage,
		immune:         immune,
		attackImmunity: attackImmunity,
		src:            src,
	}
}

// Damage returns the amount of damage that is dealt to the player.
func (ctx *HurtContext) Damage() float64 {
	return ctx.damage
}

// SetDamage sets the amount of damage that is dealt to the player.
func (ctx *HurtContext) SetDamage(damage float64) {
	ctx.damage = damage
}

// Immune returns whether the player is immune to the damage that is dealt to them.
func (ctx *HurtContext) Immune() bool {
	return ctx.immune
}

// AttackImmunity returns the duration for which the player is immune to attacks.
func (ctx *HurtContext) AttackImmunity() time.Duration {
	return ctx.attackImmunity
}

// SetAttackImmunity sets the duration for which the player is immune to attacks.
func (ctx *HurtContext) SetAttackImmunity(immunity time.Duration) {
	ctx.attackImmunity = immunity
}

// Attacker returns the attacker of the player, if the source of the damage is an attack damage source or a
// projectile damage source. If the source is neither of these, nil is returned.
func (ctx *HurtContext) Attacker() *player.Player {
	if srcAttack, ok := ctx.src.(entity.AttackDamageSource); !ok {
		if ent, ok := srcAttack.Attacker.(*player.Player); ok {
			return ent
		}

		return nil
	} else if srcProj, ok := ctx.src.(entity.ProjectileDamageSource); ok {
		if ent, ok := srcProj.Owner.(*player.Player); ok {
			return ent
		}
	}

	return nil
}

// Source returns the source of the damage that is dealt to the player.
func (ctx *HurtContext) Source() world.DamageSource {
	return ctx.src
}
