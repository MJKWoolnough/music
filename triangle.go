package music

type Triangle struct {
	phaser
}

func NewTriangle(freq, sampleRate float64) *Triangle {
	return &Triangle{newPhaser(freq, sampleRate)}
}

func (t *Triangle) process(fs []float32) {
	for i := range fs {
		n := 2 * t.next()
		if n < 0 {
			n = -n
		}
		fs[i] = n - 1
	}
}
