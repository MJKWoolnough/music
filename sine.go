package music

import "math"

type Sine struct {
	phaser
}

func NewSine(freq Note, sampleRate float64) *Sine {
	return &Sine{newPhaser(freq, sampleRate)}
}

func (s *Sine) next() float64 {
	return math.Sin(2 * math.Pi * s.phaser.next())
}
