package guitar

import "github.com/MJKWoolnough/music"

type Guitar struct {
	player *music.Player
}

func New(p *music.Player) Guitar {
	return Guitar{player: p}
}
