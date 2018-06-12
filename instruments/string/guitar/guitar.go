package guitar // import "vimagination.zapto.org/music/instruments/string/guitar"

import "vimagination.zapto.org/music"

type Guitar struct {
	player *music.Player
}

func New(p *music.Player) Guitar {
	return Guitar{player: p}
}
