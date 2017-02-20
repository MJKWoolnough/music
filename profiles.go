package music

import "math"

type Profile func(pos float64) float64

func Steady(_ float64) float64 {
	return 1
}

func Adjust(amount float64) Profile {
	return func(_ float64) float64 {
		return amount
	}
}

func AdjustProfile(amount float64, p Profile) Profile {
	return func(pos float64) float64 {
		return amount * p(pos)
	}
}

func RampUp(pos float64) float64 {
	return pos
}

func RampDown(pos float64) float64 {
	return 1 - pos
}

func ExponentialUp(pos float64) float64 {
	return math.Pow(math.E, 6*(pos-1))
}

func ExponentialDown(pos float64) float64 {
	return math.Pow(math.E, -6*pos)
}
