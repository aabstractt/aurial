package player

import (
	"github.com/aabstractt/aurial"
	"github.com/df-mc/dragonfly/server/player"
)

var (
	quitRegistry = aurial.NewRegistry[QuitHandler]()
	joinRegistry = aurial.NewRegistry[JoinHandler]()
)

type QuitHandler interface {
	// HandleQuit handles the player quitting the server.
	HandleQuit(p *player.Player)
}

type JoinHandler interface {
	// HandleJoin handles the player joining the server.
	HandleJoin(p *player.Player)
}

func QuitRegistry() *aurial.Registry[QuitHandler] {
	return quitRegistry
}

func JoinRegistry() *aurial.Registry[JoinHandler] {
	return joinRegistry
}
