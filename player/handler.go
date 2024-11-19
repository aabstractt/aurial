package player

import (
	"github.com/aabstractt/aurial/player/context"
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/entity"
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
	"time"
)

type Handler struct {
	srv *server.Server
	p   *player.Player

	player.NopHandler
}

// Apply applies the handler to a player. This should be called to apply the handler to a player.
func Apply(srv *server.Server, p *player.Player) {
	p.Handle(&Handler{srv: srv, p: p})

	for _, h := range joinRegistry.All() {
		h.HandleJoin(p)
	}
}

// HandleBlockBreak handles a block that is being broken by a player. ctx.Cancel() may be called to cancel
// the block being broken. A pointer to a slice of the block's drops is passed, and may be altered
// to change what items will actually be dropped.
func (h *Handler) HandleBlockBreak(ectx *event.Context, pos cube.Pos, drops *[]item.Stack, xp *int) {
	ctx := context.NewBlockBreakContext(pos, *drops, *xp)

	for _, handler := range blockBreakRegistry.All() {
		handler.BlockBreak(h.p, ctx)
	}

	*drops = ctx.Drops()
	*xp = ctx.XP()

	if ctx.Cancelled() {
		ectx.Cancel()
	}
}

// HandleBlockPlace handles the player placing a specific block at a position in its world. ctx.Cancel()
// may be called to cancel the block being placed.
func (h *Handler) HandleBlockPlace(ectx *event.Context, pos cube.Pos, b world.Block) {
	ctx := context.NewBlockPlaceContext(pos, b)

	for _, handler := range blockPlaceRegistry.All() {
		handler.HandleBlockPlace(h.p, ctx)
	}

	if ctx.Cancelled() {
		ectx.Cancel()
	}
}

func (h *Handler) HandleAttackEntity(ectx *event.Context, e world.Entity, force, height *float64, critical *bool) {
	ctx := context.NewAttackEntityContext(e, *force, *height, *critical)

	for _, handler := range attackEntityRegistry.All() {
		handler.HandleAttackEntity(h.p, ctx)
	}

	*force = ctx.Force()
	*height = ctx.Height()
	*critical = ctx.Critical()

	if ctx.Cancelled() {
		ectx.Cancel()
	}
}

// HandleHurt handles the player being hurt by a damage source. ctx.Cancel() may be called to cancel the
// damage being dealt to the player. The damage dealt to the player may be changed by assigning to *damage.
func (h *Handler) HandleHurt(ectx *event.Context, damage *float64, attackImmunity *time.Duration, src world.DamageSource) {
	ctx := context.NewHurtContext(*damage, *attackImmunity, src)

	for _, handler := range hurtRegistry.All() {
		handler.HandleHurt(h.p, ctx)
	}

	*damage = ctx.Damage()
	*attackImmunity = ctx.AttackImmunity()

	if ctx.Cancelled() {
		ectx.Cancel()
	}
}

// HandleDeath handles the player dying to a particular damage cause.
func (h *Handler) HandleDeath(src world.DamageSource, keepInv *bool) {
	ctx := context.NewDeathContext(src, *keepInv)
	ctx.SetMessageCondition(func(p *player.Player) bool {
		return true
	})

	if srcAttack, ok := src.(entity.AttackDamageSource); ok {
		if attacker, ok := srcAttack.Attacker.(*player.Player); ok {
			ctx.SetKiller(attacker)
		}
	}

	for _, handler := range deathRegistry.All() {
		handler.HandleDeath(h.p, ctx)
	}

	*keepInv = ctx.KeepInventory()

	msg := ctx.Message()
	if msg == "" {
		return
	}

	msgCondition := ctx.MessageCondition()
	if msgCondition == nil {
		return
	}

	for _, t := range h.srv.Players() {
		if msgCondition(t) {
			t.Message(msg)
		}
	}
}

// HandleFoodLoss handles the player losing food. ctx.Cancel() may be called to cancel the food loss.
func (h *Handler) HandleFoodLoss(ectx *event.Context, from int, to *int) {
	ctx := context.NewFoodLossContext(from, *to)

	for _, handler := range foodLossRegistry.All() {
		handler.HandleFoodLoss(h.p, ctx)
	}

	*to = ctx.To()

	if ctx.Cancelled() {
		ectx.Cancel()
	}
}

// HandleHeal handles the player being healed by a healing source. ctx.Cancel() may be called to cancel
// the healing.
func (h *Handler) HandleHeal(ectx *event.Context, health *float64, src world.HealingSource) {
	ctx := context.NewHealContext(*health, src)

	for _, handler := range healRegistry.All() {
		handler.HandleHeal(h.p, ctx)
	}

	*health = ctx.Amount()

	if ctx.Cancelled() {
		ectx.Cancel()
	}
}

// HandleChat handles a player sending a chat message. ctx.Cancel() may be called to cancel the chat message.
func (h *Handler) HandleChat(ectx *event.Context, message *string) {
	ctx := context.NewChatContext(*message)

	for _, handler := range chatRegistry.All() {
		handler.HandleChat(h.p, ctx)
	}

	*message = ctx.Message()

	if ctx.Cancelled() {
		ectx.Cancel()
	}
}

// HandleChangeWorld handles when the player is added to a new world. before may be nil.
func (h *Handler) HandleChangeWorld(before, after *world.World) {
	ctx := context.NewChangeWorldContext(before, after)

	for _, handler := range changeWorldRegistry.All() {
		handler.HandleChangeWorld(h.p, ctx)
	}

	after = ctx.After()
}

// HandleTeleport handles the teleportation of a player. ctx.Cancel() may be called to cancel it.
func (h *Handler) HandleTeleport(ectx *event.Context, pos mgl64.Vec3) {
	ctx := context.NewTeleportContext(pos)

	for _, handler := range teleportRegistry.All() {
		handler.HandleTeleport(h.p, ctx)
	}

	if ctx.Cancelled() {
		ectx.Cancel()
	}
}

// HandleMove handles the movement of a player. ctx.Cancel() may be called to cancel the movement event.
// The new position, yaw and pitch are passed.
func (h *Handler) HandleMove(ectx *event.Context, newPos mgl64.Vec3, newYaw, newPitch float64) {
	ctx := context.NewMoveContext(newPos, newYaw, newPitch)

	for _, handler := range moveRegistry.All() {
		handler.HandleMove(h.p, ctx)
	}

	if ctx.Cancelled() {
		ectx.Cancel()
	}
}

// HandleQuit handles the player quitting the game.
func (h *Handler) HandleQuit() {
	for _, handler := range quitRegistry.All() {
		handler.HandleQuit(h.p)
	}
}

// HandleItemUse handles the player using an item. ctx.Cancel() may be called to cancel the item use.
func (h *Handler) HandleItemUse(ectx *event.Context) {
	ctx := context.NewItemUseContext(h.p.HeldItems())

	for _, handler := range itemUseRegistry.All() {
		handler.HandleItemUse(h.p, ctx)
	}

	if ctx.Cancelled() {
		ectx.Cancel()
	}
}
