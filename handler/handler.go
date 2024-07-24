package handler

import (
    "github.com/df-mc/dragonfly/server/block/cube"
    "github.com/df-mc/dragonfly/server/event"
    "github.com/df-mc/dragonfly/server/item"
    "github.com/df-mc/dragonfly/server/player"
    "github.com/df-mc/dragonfly/server/world"
    "github.com/go-gl/mathgl/mgl64"
    "sync"
    "time"
)

const (
    BlockBreakHandlerID = iota
    BlockPlaceHandlerID
    AttackEntityHandlerID
    HurtHandlerID
    DeathHandlerID
    FoodLossHandlerID
    HealHandlerID
    ChatHandlerID
    ChangeWorldHandlerID
    TeleportHandlerID
    MoveHandlerID
)

var (
    handlers   = make(map[int][]interface{})
    handlersMu sync.Mutex
)

type Handler struct {
    p *player.Player

    player.NopHandler
}

// Hook hooks the handler into the player p. This should be called when a player is created.
func Hook(p *player.Player) {
    p.Handle(&Handler{p: p})
}

// RegisterHandler registers a handler for a specific event ID. The handler will be called when the event
func RegisterHandler(id int, handler interface{}) {
    handlersMu.Lock()
    defer handlersMu.Unlock()

    handlers[id] = append(handlers[id], handler)
}

// HandleBlockBreak handles a block that is being broken by a player. ctx.Cancel() may be called to cancel
// the block being broken. A pointer to a slice of the block's drops is passed, and may be altered
// to change what items will actually be dropped.
func (h *Handler) HandleBlockBreak(ctx *event.Context, pos cube.Pos, drops *[]item.Stack, xp *int) {
    handlersMu.Lock()
    defer handlersMu.Unlock()

    for _, handler := range handlers[BlockBreakHandlerID] {
        if handler.(BlockBreakHandler).BlockBreak(h.p, pos, drops, xp, ctx.Cancelled()) {
            ctx.Cancel()
        }
    }
}

// HandleBlockPlace handles the player placing a specific block at a position in its world. ctx.Cancel()
// may be called to cancel the block being placed.
func (h *Handler) HandleBlockPlace(ctx *event.Context, pos cube.Pos, b world.Block) {
    handlersMu.Lock()
    defer handlersMu.Unlock()

    for _, handler := range handlers[BlockPlaceHandlerID] {
        if handler.(BlockPlaceHandler).BlockPlace(h.p, pos, b, ctx.Cancelled()) {
            ctx.Cancel()
        }
    }
}

func (h *Handler) HandleAttackEntity(ctx *event.Context, e world.Entity, force, height *float64, critical *bool) {
    handlersMu.Lock()
    defer handlersMu.Unlock()

    for _, handler := range handlers[AttackEntityHandlerID] {
        if handler.(AttackEntityHandler).HandleAttackEntity(h.p, e, force, height, critical, ctx.Cancelled()) {
            ctx.Cancel()
        }
    }
}

// HandleHurt handles the player being hurt by a damage source. ctx.Cancel() may be called to cancel the
// damage being dealt to the player. The damage dealt to the player may be changed by assigning to *damage.
func (h *Handler) HandleHurt(ctx *event.Context, damage *float64, attackImmunity *time.Duration, src world.DamageSource) {
    handlersMu.Lock()
    defer handlersMu.Unlock()

    for _, handler := range handlers[HurtHandlerID] {
        if handler.(HurtHandler).HandleHurt(h.p, damage, attackImmunity, src, ctx.Cancelled()) {
            ctx.Cancel()
        }
    }
}

// HandleDeath handles the player dying to a particular damage cause.
func (h *Handler) HandleDeath(src world.DamageSource, keepInv *bool) {
    handlersMu.Lock()
    defer handlersMu.Unlock()

    for _, handler := range handlers[DeathHandlerID] {
        handler.(DeathHandler).HandleDeath(h.p, src, keepInv)
    }
}

// HandleFoodLoss handles the player losing food. ctx.Cancel() may be called to cancel the food loss.
func (h *Handler) HandleFoodLoss(ctx *event.Context, from int, to *int) {
    handlersMu.Lock()
    defer handlersMu.Unlock()

    for _, handler := range handlers[FoodLossHandlerID] {
        if handler.(FoodLossHandler).HandleFoodLoss(h.p, from, to, ctx.Cancelled()) {
            ctx.Cancel()
        }
    }
}

// HandleHeal handles the player being healed by a healing source. ctx.Cancel() may be called to cancel
// the healing.
func (h *Handler) HandleHeal(ctx *event.Context, health *float64, src world.HealingSource) {
    handlersMu.Lock()
    defer handlersMu.Unlock()

    for _, handler := range handlers[HealHandlerID] {
        if handler.(HealHandler).HandleHeal(h.p, health, src, ctx.Cancelled()) {
            ctx.Cancel()
        }
    }
}

// HandleChat handles a player sending a chat message. ctx.Cancel() may be called to cancel the chat message.
func (h *Handler) HandleChat(ctx *event.Context, message *string) {
    handlersMu.Lock()
    defer handlersMu.Unlock()

    for _, handler := range handlers[ChatHandlerID] {
        if handler.(ChatHandler).HandleChat(h.p, message, ctx.Cancelled()) {
            ctx.Cancel()
        }
    }
}

// HandleChangeWorld handles when the player is added to a new world. before may be nil.
func (h *Handler) HandleChangeWorld(before, after *world.World) {
    handlersMu.Lock()
    defer handlersMu.Unlock()

    for _, handler := range handlers[ChangeWorldHandlerID] {
        handler.(ChangeWorldHandler).HandleChangeWorld(h.p, before, after)
    }
}

// HandleTeleport handles the teleportation of a player. ctx.Cancel() may be called to cancel it.
func (h *Handler) HandleTeleport(ctx *event.Context, pos mgl64.Vec3) {
    handlersMu.Lock()
    defer handlersMu.Unlock()

    for _, handler := range handlers[TeleportHandlerID] {
        if handler.(TeleportHandler).HandleTeleport(h.p, pos, ctx.Cancelled()) {
            ctx.Cancel()
        }
    }
}

// HandleMove handles the movement of a player. ctx.Cancel() may be called to cancel the movement event.
// The new position, yaw and pitch are passed.
func (h *Handler) HandleMove(ctx *event.Context, newPos mgl64.Vec3, newYaw, newPitch float64) {
    handlersMu.Lock()
    defer handlersMu.Unlock()

    for _, handler := range handlers[MoveHandlerID] {
        if handler.(MoveHandler).HandleMove(h.p, newPos, newYaw, newPitch, ctx.Cancelled()) {
            ctx.Cancel()
        }
    }
}
