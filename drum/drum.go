package drum

import "github.com/MJKWoolnough/music"

type Drum struct {
	player *music.Player
}

func New(p *music.Player) Drum {
	return Drum{player: p}
}
