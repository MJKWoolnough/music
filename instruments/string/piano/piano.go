package piano

import "github.com/MJKWoolnough/music"

type Piano struct {
	player *music.Player
}

func New(p *music.Player) Piano {
	return Piano{player: p}
}
