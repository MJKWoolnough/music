package trombone

import "github.com/MJKWoolnough/music"

type Trombone struct {
	player *music.Player
}

func New(p *music.Player) Trombone {
	return Trombone{player: p}
}
