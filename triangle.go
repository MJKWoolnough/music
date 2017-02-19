package music

func Triangle(phase float64) float64 {
	if phase < 0.5 {
		return 4*phase - 1
	}
	return 2 - 4*phase
}
