package handler

import (
    "github.com/df-mc/dragonfly/server/player"
    "github.com/df-mc/dragonfly/server/world"
)

type ChangeWorldHandler interface {
    // HandleChangeWorld handles when the player is added to a new world. before may be nil.
    HandleChangeWorld(p *player.Player, before, after *world.World)
}
