package music

type Triangle struct {
	phaser
}

func NewTriangle(freq Note, sampleRate float64) *Triangle {
	return &Triangle{newPhaser(freq, sampleRate)}
}

func (t *Triangle) next() float64 {
	n := 2 * t.phaser.next()
	if n < 0 {
		n = -n
	}
	return n - 1
}
