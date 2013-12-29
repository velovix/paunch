package paunch

// Point is an object that represents an X and Y position in 2D space.
type Point struct {
	X float64
	Y float64
}

func getPointDistance(point1, point2 *Point) *Point {

	distance := NewPoint(point2.X-point1.X, point2.Y-point1.Y)

	return distance
}

// NewPoint creates a new Point object.
func NewPoint(x, y float64) *Point {

	return &Point{X: x, Y: y}
}

// Move moves the Point object a specified distance.
func (point *Point) Move(x, y float64) {

	point.X += x
	point.Y += y
}

// DistanceToTangentPoint returns a Point with values equal to the distance
// a given Point is from the closest tangent Point on the given side of the
// Point. However, for a Point, this doesn't mean much. All this method returns
// is the distance between a given Point and the Point itself.
func (point *Point) DistanceToTangentPoint(point2 *Point, side Direction) *Point {

	return getPointDistance(point, point2)
}

// OnPoint checks if a Point is on the Point object.
func (point *Point) OnPoint(point2 *Point) bool {

	if point.X == point2.X && point.Y == point2.Y {
		return true
	}

	return false
}
