package player

import (
	"github.com/aabstractt/aurial"
	"github.com/df-mc/dragonfly/server/player"
)

var quitRegistry = aurial.NewRegistry[QuitHandler]()

type QuitHandler interface {
	// HandleQuit handles the player quitting the server.
	HandleQuit(p *player.Player)
}

func QuitRegistry() *aurial.Registry[QuitHandler] {
	return quitRegistry
}
