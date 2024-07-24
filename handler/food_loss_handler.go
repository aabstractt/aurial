package handler

import (
    "github.com/df-mc/dragonfly/server/player"
)

type FoodLossHandler interface {
    // HandleFoodLoss handles the food bar of a player depleting naturally, for example because the player was
    // sprinting and jumping. If the food loss should be cancelled, return true.
    // cancelled is true if the food loss already was cancelled by another handler.
    HandleFoodLoss(p *player.Player, from int, to *int, cancelled bool) bool
}
