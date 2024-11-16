package context

import (
	"github.com/df-mc/dragonfly/server/world"
	"time"
)

type HurtContext struct {
	// Damage is the amount of damage that is dealt to the player. This value may be changed to change the
	// amount of damage dealt to the player.
	damage float64

	// AttackImmunity is the duration for which the player is immune to attacks. This value may be changed to
	// change the duration of immunity.
	attackImmunity time.Duration

	// Source is the source of the damage that is dealt to the player.
	src world.DamageSource

	CancellableContext
}

// Damage returns the amount of damage that is dealt to the player.
func (h *HurtContext) Damage() float64 {
	return h.damage
}

// SetDamage sets the amount of damage that is dealt to the player.
func (h *HurtContext) SetDamage(damage float64) {
	h.damage = damage
}

// AttackImmunity returns the duration for which the player is immune to attacks.
func (h *HurtContext) AttackImmunity() time.Duration {
	return h.attackImmunity
}

// SetAttackImmunity sets the duration for which the player is immune to attacks.
func (h *HurtContext) SetAttackImmunity(immunity time.Duration) {
	h.attackImmunity = immunity
}

// Source returns the source of the damage that is dealt to the player.
func (h *HurtContext) Source() world.DamageSource {
	return h.src
}
