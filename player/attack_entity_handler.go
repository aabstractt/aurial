package player

import (
	"github.com/aabstractt/aurial"
	"github.com/aabstractt/aurial/player/context"
	"github.com/df-mc/dragonfly/server/player"
)

var attackEntityRegistry = aurial.NewRegistry[AttackEntityHandler]()

type AttackEntityHandler interface {
	// HandleAttackEntity is called when a player attacks an entity. The entity that is being attacked is
	// passed in the context.
	HandleAttackEntity(p *player.Player, ctx *context.AttackEntityContext)
}

func AttackEntityRegistry() *aurial.Registry[AttackEntityHandler] {
	return attackEntityRegistry
}
