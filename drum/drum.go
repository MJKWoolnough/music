package drum

import "github.com/MJKWoolnough/music"

type Drum struct {
	player *music.Player
}

func New(p *music.Player) Drum {
	return Drum{player: p}
}

func (d Drum) Bass(time uint64, channel int) {
	d.player.Add(time, d.player.Rate()*2, music.E2, music.Sine, music.ExponentialDown, channel)
}
