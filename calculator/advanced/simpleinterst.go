package advanced

func Calculate(p float64, r float64, t float64) (intrest float64) {
	intrest = p * r * float64(t) / 100
	return
}
