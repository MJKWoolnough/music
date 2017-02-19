package music

func Steady(_ float64) float64 {
	return 1
}

func Adjust(amount_ float64) func(float64) float64 {
	return func(_ float64) float64 {
		return amount
	}
}

func RampUp(pos float64) float64 {
	return pos
}

func RampDown(pos float64) float64 {
	return 1 - pos
}
