package world

import (
	"github.com/aabstractt/aurial/world/context"
	"github.com/df-mc/dragonfly/server/world"
)

type handler struct {
	world.NopHandler
}

func Apply(w *world.World) {
	w.Handle(&handler{})
}

func (h *handler) HandleEntitySpawn(tx *world.Tx, e world.Entity) {
	ctx := context.NewEntitySpawnContext(e)

	for _, hand := range entitySpawnRegistry.All() {
		hand.HandleEntitySpawn(tx, ctx)
	}
}
