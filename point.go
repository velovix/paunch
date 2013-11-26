package paunch

// Point is an object that represents an X and Y position in 2D space.
type Point struct {
	X float64
	Y float64
}

// NewPoint creates a new Point object.
func NewPoint(x, y float64) Point {

	return Point{X: x, Y: y}
}

// OnPoint checks if a Point is on the Point object.
func (point1 Point) OnPoint(point2 Point) bool {

	if point1.X == point2.X && point1.Y == point2.Y {
		return true
	}

	return false
}

// GetSlope returns the slope created by two points.
func GetSlope(point1, point2 Point) float64 {

	return (point2.Y - point1.Y) / (point2.X - point1.X)
}
