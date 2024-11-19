package context

import "github.com/df-mc/dragonfly/server/world"

type EntitySpawnContext struct {
	entity world.Entity
}

// NewEntitySpawnContext returns a new EntitySpawnContext with the entity that is being spawned set.
func NewEntitySpawnContext(entity world.Entity) *EntitySpawnContext {
	return &EntitySpawnContext{entity: entity}
}

// Entity returns the entity that is being spawned.
func (ctx *EntitySpawnContext) Entity() world.Entity {
	return ctx.entity
}
