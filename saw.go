package music

type Saw struct {
	phaser
}

func NewSaw(freq Note, sampleRate float64) *Saw {
	return &Saw{newPhaser(freq, sampleRate)}
}
