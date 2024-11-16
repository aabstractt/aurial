package context

import "github.com/df-mc/dragonfly/server/world"

type AttackEntityContext struct {
	e world.Entity

	force  float64
	height float64

	critical bool

	CancellableContext
}

// Entity returns the entity that is being attacked.
func (ctx *AttackEntityContext) Entity() world.Entity {
	return ctx.e
}

// Force returns the force of the movement.
func (ctx *AttackEntityContext) Force() float64 {
	return ctx.force
}

// SetForce sets the force of the movement.
func (ctx *AttackEntityContext) SetForce(force float64) {
	ctx.force = force
}

// Height returns the height of the movement.
func (ctx *AttackEntityContext) Height() float64 {
	return ctx.height
}

// SetHeight sets the height of the movement.
func (ctx *AttackEntityContext) SetHeight(height float64) {
	ctx.height = height
}

// Critical returns whether the attack is critical.
func (ctx *AttackEntityContext) Critical() bool {
	return ctx.critical
}

// SetCritical sets whether the attack is critical.
func (ctx *AttackEntityContext) SetCritical(critical bool) {
	ctx.critical = critical
}

func NewAttackEntityContext(e world.Entity, force, height float64, critical bool) *AttackEntityContext {
	return &AttackEntityContext{
		e: e,

		force:  force,
		height: height,

		critical: critical,
	}
}
