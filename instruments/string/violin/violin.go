package violin // import "vimagination.zapto.org/music/instruments/string/violin"

import "vimagination.zapto.org/music"

type Violin struct {
	player *music.Player
}

func New(p *music.Player) Violin {
	return Violin{player: p}
}
