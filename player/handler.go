package player

import (
	"github.com/aabstractt/aurial/player/context"
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/entity"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
	"time"
)

type Handler struct {
	srv *server.Server

	player.NopHandler
}

// Apply applies the handler to a player. This should be called to apply the handler to a player.
func Apply(srv *server.Server, p *player.Player) {
	p.Handle(&Handler{srv: srv})

	for _, h := range joinRegistry.All() {
		h.HandleJoin(p)
	}
}

// HandleBlockBreak handles a block that is being broken by a player. ctx.Cancel() may be called to cancel
// the block being broken. A pointer to a slice of the block's drops is passed, and may be altered
// to change what items will actually be dropped.
func (h *Handler) HandleBlockBreak(ectx *player.Context, pos cube.Pos, drops *[]item.Stack, xp *int) {
	ctx := context.NewBlockBreakContext(pos, *drops, *xp)

	for _, handler := range blockBreakRegistry.All() {
		handler.BlockBreak(ectx.Val(), ctx)
	}

	*drops = ctx.Drops()
	*xp = ctx.XP()

	if ctx.Cancelled() {
		ectx.Cancel()
	}
}

// HandleBlockPlace handles the player placing a specific block at a position in its world. ctx.Cancel()
// may be called to cancel the block being placed.
func (h *Handler) HandleBlockPlace(ectx *player.Context, pos cube.Pos, b world.Block) {
	ctx := context.NewBlockPlaceContext(pos, b)

	for _, handler := range blockPlaceRegistry.All() {
		handler.HandleBlockPlace(ectx.Val(), ctx)
	}

	if ctx.Cancelled() {
		ectx.Cancel()
	}
}

func (h *Handler) HandleAttackEntity(ectx *player.Context, e world.Entity, force, height *float64, critical *bool) {
	ctx := context.NewAttackEntityContext(e, *force, *height, *critical)

	for _, handler := range attackEntityRegistry.All() {
		handler.HandleAttackEntity(ectx.Val(), ctx)
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
func (h *Handler) HandleHurt(ectx *player.Context, damage *float64, immune bool, attackImmunity *time.Duration, src world.DamageSource) {
	ctx := context.NewHurtContext(*damage, immune, *attackImmunity, src)

	for _, handler := range hurtRegistry.All() {
		handler.HandleHurt(ectx.Val(), ctx)
	}

	*damage = ctx.Damage()
	*attackImmunity = ctx.AttackImmunity()

	if ctx.Cancelled() {
		ectx.Cancel()
	}
}

// HandleDeath handles the player dying to a particular damage cause.
func (h *Handler) HandleDeath(p *player.Player, src world.DamageSource, keepInv *bool) {
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
		handler.HandleDeath(p, ctx)
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

	for t := range h.srv.Players(nil) {
		if msgCondition(t) {
			t.Message(msg)
		}
	}
}

// HandleFoodLoss handles the player losing food. ctx.Cancel() may be called to cancel the food loss.
func (h *Handler) HandleFoodLoss(ectx *player.Context, from int, to *int) {
	ctx := context.NewFoodLossContext(from, *to)

	for _, handler := range foodLossRegistry.All() {
		handler.HandleFoodLoss(ectx.Val(), ctx)
	}

	*to = ctx.To()

	if ctx.Cancelled() {
		ectx.Cancel()
	}
}

// HandleHeal handles the player being healed by a healing source. ctx.Cancel() may be called to cancel
// the healing.
func (h *Handler) HandleHeal(ectx *player.Context, health *float64, src world.HealingSource) {
	ctx := context.NewHealContext(*health, src)

	for _, handler := range healRegistry.All() {
		handler.HandleHeal(ectx.Val(), ctx)
	}

	*health = ctx.Amount()

	if ctx.Cancelled() {
		ectx.Cancel()
	}
}

// HandleChat handles a player sending a chat message. ctx.Cancel() may be called to cancel the chat message.
func (h *Handler) HandleChat(ectx *player.Context, message *string) {
	ctx := context.NewChatContext(*message)

	for _, handler := range chatRegistry.All() {
		handler.HandleChat(ectx.Val(), ctx)
	}

	*message = ctx.Message()

	if ctx.Cancelled() {
		ectx.Cancel()
	}
}

// HandleChangeWorld handles when the player is added to a new world. before may be nil.
func (h *Handler) HandleChangeWorld(p *player.Player, before, after *world.World) {
	ctx := context.NewChangeWorldContext(before, after)

	for _, handler := range changeWorldRegistry.All() {
		handler.HandleChangeWorld(p, ctx)
	}

	after = ctx.After()
}

// HandleTeleport handles the teleportation of a player. ctx.Cancel() may be called to cancel it.
func (h *Handler) HandleTeleport(ectx *player.Context, pos mgl64.Vec3) {
	ctx := context.NewTeleportContext(pos)

	for _, handler := range teleportRegistry.All() {
		handler.HandleTeleport(ectx.Val(), ctx)
	}

	if ctx.Cancelled() {
		ectx.Cancel()
	}
}

// HandleMove handles the movement of a player. ctx.Cancel() may be called to cancel the movement event.
// The new position, yaw and pitch are passed.
func (h *Handler) HandleMove(ectx *player.Context, newPos mgl64.Vec3, rot cube.Rotation) {
	ctx := context.NewMoveContext(newPos, rot)

	for _, handler := range moveRegistry.All() {
		handler.HandleMove(ectx.Val(), ctx)
	}

	if ctx.Cancelled() {
		ectx.Cancel()
	}
}

// HandleQuit handles the player quitting the game.
func (h *Handler) HandleQuit(p *player.Player) {
	for _, handler := range quitRegistry.All() {
		handler.HandleQuit(p)
	}
}

// HandleItemUse handles the player using an item. ctx.Cancel() may be called to cancel the item use.
func (h *Handler) HandleItemUse(ectx *player.Context) {
	p := ectx.Val()
	ctx := context.NewItemUseContext(p.HeldItems())

	for _, handler := range itemUseRegistry.All() {
		handler.HandleItemUse(p, ctx)
	}

	if ctx.Cancelled() {
		ectx.Cancel()
	}
}

func (h *Handler) HandlePunchAir(ctx *player.Context) {
	for _, handler := range punchAirRegistry.All() {
		handler.HandlePunchAir(ctx.Val())
	}
}
