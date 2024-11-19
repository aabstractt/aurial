package player

import (
	"github.com/aabstractt/aurial"
	"github.com/aabstractt/aurial/player/context"
	"github.com/df-mc/dragonfly/server/player"
)

var foodLossRegistry = aurial.NewRegistry[FoodLossHandler]()

type FoodLossHandler interface {
	// HandleFoodLoss handles the food bar of a player depleting naturally, for example because the player was
	// sprinting and jumping. If the food loss should be cancelled use {ctx.Cancel()}.
	HandleFoodLoss(p *player.Player, ctx *context.FoodLossContext)
}

func FoodLossRegistry() *aurial.Registry[FoodLossHandler] {
	return foodLossRegistry
}
