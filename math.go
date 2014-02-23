package paunch

import (
	"math"
)

func findDeterminate(a, b, c, d float64) float64 {

	return (a * d) - (b * c)
}

func getSlope(x1, y1, x2, y2 float64) float64 {

	return (y2 - y1) / (x2 - x1)
}

func round(x float64, prec int) float64 {

	var rounder float64

	pow := math.Pow(10, float64(prec))
	intermed := x * pow
	_, frac := math.Modf(intermed)

	if frac >= 0.5 {
		rounder = math.Ceil(intermed)
	} else {
		rounder = math.Floor(intermed)
	}

	return rounder / pow
}
