package music

import "math"

func Sine(phase float64) float64 {
	return math.Sin(2 * math.Pi * phase)
}
