package music

func Beat(p *Player, note Note, start, length uint64, channel int) {
	tu := p.Rate() / 100
	lu := p.EndZeroNote(p.Rate()/100, note)
	ld := p.EndZeroNote(length-tu, note)
	p.Add(start, lu, note, Sine, ExponentialUp, channel)
	p.Add(start+lu, ld, note, Sine, ExponentialDown, channel)
}
