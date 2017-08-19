package violin

import "github.com/MJKWoolnough/music"

type Violin struct {
	player *music.Player
}

func New(p *music.Player) Violin {
	return Violin{player: p}
}
