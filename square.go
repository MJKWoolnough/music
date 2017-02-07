package music

type Square struct {
	phaser
}

func NewSquare(freq Note, sampleRate float64) *Square {
	return &Square{newPhaser(freq, sampleRate)}
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
