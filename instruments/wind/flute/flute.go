package flute

import "github.com/MJKWoolnough/music"

type Flute struct {
	player *music.Player
}

func New(p *music.Player) Flute {
	return Flute{player: p}
}
