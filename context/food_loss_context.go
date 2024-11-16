package context

type FoodLossContext struct {
	from, to int
}

// From returns the amount of food the player had before the food loss.
func (ctx *FoodLossContext) From() int {
	return ctx.from
}

// To returns the amount of food the player has after the food loss.
func (ctx *FoodLossContext) To() int {
	return ctx.to
}

// SetTo sets the amount of food the player has after the food loss.
func (ctx *FoodLossContext) SetTo(to int) {
	ctx.to = to
}
