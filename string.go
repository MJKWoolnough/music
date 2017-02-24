package music

import "math/rand"

// https://uk.mathworks.com/help/signal/examples/generating-guitar-chords-using-the-karplus-strong-algorithm.html
// http://www.cs.princeton.edu/courses/archive/fall07/cos126/assignments/guitar.html

type String struct {
	data  []float64
	pos   int
	decay float64
}

const StringDecay = 0.4998

func NewString(n Note, rate, decay float64) *String {
	return &String{
		data:  make([]float64, int(rate/float64(n))),
		decay: decay,
	}
}

func (s *String) Pluck() {
	for i := range s.data {
		s.data[i] = rand.Float64()*2 - 1
	}
	s.pos = 0
}

func (s *String) Next(_ float64) float64 {
	p := s.pos
	f := s.data[p]
	s.pos++
	if s.pos == len(s.data) {
		s.pos = 0
	}
	s.data[p] = s.decay * (f + s.data[s.pos])
	return f
}
