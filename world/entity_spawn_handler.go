package world

import (
	"github.com/aabstractt/aurial"
	"github.com/aabstractt/aurial/world/context"
	"github.com/df-mc/dragonfly/server/world"
)

var entitySpawnRegistry = aurial.NewRegistry[EntitySpawnHandler]()

type EntitySpawnHandler interface {
	// HandleEntitySpawn handles the spawning of an entity in the world.
	HandleEntitySpawn(tx *world.Tx, ctx *context.EntitySpawnContext)
}

func EntitySpawnRegistry() *aurial.Registry[EntitySpawnHandler] {
	return entitySpawnRegistry
}
