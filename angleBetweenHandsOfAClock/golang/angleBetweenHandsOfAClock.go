package angleBetweenHandsOfAClock

func angleClock(hour int, minutes int) float64 {
	ah := float64(hour*30) + float64(minutes*30)/float64(60)
	am := float64(minutes * 6)
	a := ah - am
	if a < 0 {
		a = -a
	}

	if a > 180 {
		a = 360 - a
	}
	return a
}
