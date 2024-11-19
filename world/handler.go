package world

import (
	"github.com/aabstractt/aurial/world/context"
	"github.com/df-mc/dragonfly/server/world"
)

type handler struct {
	world.NopHandler

	w *world.World
}

func Apply(w *world.World) {
	w.Handle(&handler{w: w})
}

func (h *handler) HandleEntitySpawn(e world.Entity) {
	ctx := context.NewEntitySpawnContext(e)

	for _, hand := range entitySpawnRegistry.All() {
		hand.HandleEntitySpawn(h.w, ctx)
	}
}
