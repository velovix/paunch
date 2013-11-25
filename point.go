package paunch

type Point struct {
	X float64
	Y float64
}

func NewPoint(x, y float64) Point {

	return Point{X: x, Y: y}
}

func isPointOnPoint(point1, point2 Point) bool {

	if point1.X == point2.X && point1.Y == point2.Y {
		return true
	}

	return false
}

func GetSlope(point1, point2 Point) float64 {

	return (point2.Y - point1.Y) / (point2.X - point1.X)
}
