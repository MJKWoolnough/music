package piano // import "vimagination.zapto.org/music/instruments/string/piano"

import "vimagination.zapto.org/music"

type Piano struct {
	player *music.Player
}

func New(p *music.Player) Piano {
	return Piano{player: p}
}
