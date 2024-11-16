package aurial

import (
	"github.com/aabstractt/aurial/context"
	"github.com/df-mc/dragonfly/server/player"
)

type AttackEntityHandler interface {
	// HandleAttackEntity is called when a player attacks an entity. The entity that is being attacked is
	// passed in the context.
	HandleAttackEntity(p *player.Player, ctx *context.AttackEntityContext)
}
