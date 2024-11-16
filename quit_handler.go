package aurial

import "github.com/df-mc/dragonfly/server/player"

type QuitHandler interface {
	// HandleQuit handles the player quitting the server.
	HandleQuit(p *player.Player)
}
