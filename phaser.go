package music

type phaser struct {
	phase, speed float64
}

func newPhaser(freq, rate float64) phaser {
	return phaser{speed: freq / rate}
}

func (p *phaser) next() float64 {
	_, p.phase += p.speed
	for p.phase > 1 {
		p.phase -= 1
	}
	for p.phase < -1 {
		p.phase += 1
	}
	return p.phase
}
