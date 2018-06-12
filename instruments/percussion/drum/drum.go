package drum // import "vimagination.zapto.org/music/instruments/percussion/drum"

import "vimagination.zapto.org/music"

type Drum struct {
	player *music.Player
}

func New(p *music.Player) Drum {
	return Drum{player: p}
}

func (d Drum) Bass(time uint64, channel int) {
	rate := d.player.Rate()
	bassIntro := d.player.EndZeroNote(rate/50, music.E2)
	d.player.Add(time, d.player.EndZeroNote(rate/50, music.E3), music.E3.Sine, music.RampUp, channel)
	d.player.Add(time, d.player.EndZeroNote(rate/75, music.E4), music.E4.Sine, music.AdjustProfile(0.75, music.RampUp), channel)
	d.player.Add(time, d.player.EndZeroNote(rate/100, music.E5), music.E5.Sine, music.AdjustProfile(0.5, music.RampUp), channel)
	d.player.Add(time+bassIntro, rate*2, music.E2.Sine, music.ExponentialDown, channel)
}
