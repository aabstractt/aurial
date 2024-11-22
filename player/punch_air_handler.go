package player

import (
	"github.com/aabstractt/aurial"
	"github.com/df-mc/dragonfly/server/player"
)

var punchAirRegistry = aurial.NewRegistry[PunchAirHandler]()

type PunchAirHandler interface {
	// HandlePunchAir is called when a player punches air.
	HandlePunchAir(p *player.Player)
}

func PunchAirRegistry() *aurial.Registry[PunchAirHandler] {
	return punchAirRegistry
}
