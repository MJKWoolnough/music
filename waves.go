package music

import "math"

type Wave func(phase float64) float64

func (n Note) Saw(phase float64) float64 {
	_, frac := math.Modf(float64(n) * phase)
	return frac*2 - 1
}

func (n Note) Sine(phase float64) float64 {
	_, frac := math.Modf(float64(n) * phase)
	return math.Sin(2 * math.Pi * frac)
}
func (n Note) Square(phase float64) float64 {
	_, frac := math.Modf(float64(n) * phase)
	if frac > 0.5 {
		return 1
	}
	return -1
}

func (n Note) Triangle(phase float64) float64 {
	_, frac := math.Modf(float64(n) * phase)
	if frac < 0.5 {
		return 4*frac - 1
	}
	return 3 - 4*frac
}
