package music

type Square struct {
	phaser
}

func NewSquare(freq Note, sampleRate float64) *Square {
	return &Square{newPhaser(freq, sampleRate)}
}

func (s *Square) next() float64 {
	if s.phaser.next() > 0.5 {
		return 1
	}
	return -1
}
