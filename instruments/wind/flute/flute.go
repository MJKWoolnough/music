package flute // import "vimagination.zapto.org/music/instruments/wind/flute"

import "vimagination.zapto.org/music"

type Flute struct {
	player *music.Player
}

func New(p *music.Player) Flute {
	return Flute{player: p}
}
