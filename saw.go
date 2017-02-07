package music

type Saw struct {
	phaser
}

func NewSaw(freq, sampleRate float64) *Saw {
	return &Saw{newPhaser(freq, sampleRate, 0)}
}

func (s *Saw) process(fs []float32) {
	for i := range fs {
		fs[i] = float32(s.next())
	}
}
