package music

type Square struct {
	phaser
}

func NewSquare(freq, sampleRate float64) *Square {
	return &Square{newPhaser(freq, sampleRate, 0)}
}

func (s *Square) process(fs []float32) {
	for i := range fs {
		if s.next() > 0.5 {
			fs[i] = 1
		} else {
			fs[i] = -1
		}
	}
}
