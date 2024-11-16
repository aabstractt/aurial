package context

import "github.com/df-mc/dragonfly/server/world"

type HealContext struct {
	amount float64 // amount is the amount of health that will be healed.

	src world.HealingSource // HealingSource is the source of the healing.

	CancellableContext
}

// Amount returns the amount of health that will be healed.
func (ctx *HealContext) Amount() float64 {
	return ctx.amount
}

// SetAmount sets the amount of health that will be healed.
func (ctx *HealContext) SetAmount(amount float64) {
	ctx.amount = amount
}

// Source returns the source of the healing.
func (ctx *HealContext) Source() world.HealingSource {
	return ctx.src
}
