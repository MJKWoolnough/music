package trombone // import "vimagination.zapto.org/music/instruments/brass/trombone"

import "vimagination.zapto.org/music"

type Trombone struct {
	player *music.Player
}

func New(p *music.Player) Trombone {
	return Trombone{player: p}
}
