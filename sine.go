package music

import "math"

type Sine struct {
	phraser
}

func NewSine(freq, sampleRate float64) *Sine {
	return &Sine{newPhaser(freq, sampleRate)}
}

func (s *Sine) process(fs []float32) {
	for i := range fs {
		fs[i] = float32(math.Sin(2 * math.Pi * s.next()))
	}
}
