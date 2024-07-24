package handler

import (
    "github.com/df-mc/dragonfly/server/event"
    "github.com/df-mc/dragonfly/server/player"
)

type FoodLossHandler interface {
    // HandleFoodLoss handles the food bar of a player depleting naturally, for example because the player was
    // sprinting and jumping. ctx.Cancel() may be called to cancel the food points being lost.
    HandleFoodLoss(p *player.Player, ctx *event.Context, from int, to *int)
}
