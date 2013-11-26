package paunch

func findDeterminate(a, b, c, d float64) float64 {

	return (a * d) - (b * c)
}

func getSlope(x1, y1, x2, y2 float64) float64 {

	return (y2 - y1) / (x2 - x1)
}
