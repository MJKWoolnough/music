package music

import "math"

type Wave func(phase float64) float64

func Saw(phase float64) float64 {
	return phase*2 - 1
}

func Sine(phase float64) float64 {
	return math.Sin(2 * math.Pi * phase)
}
func Square(phase float64) float64 {
	if phase > 0.5 {
		return 1
	}
	return -1
}

func Triangle(phase float64) float64 {
	if phase < 0.5 {
		return 4*phase - 1
	}
	return 2 - 4*phase
}
