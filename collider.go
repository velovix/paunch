package paunch

// Collider is an interface that qualifies types that can check for collisions
// against Points, Boundings, Lines, and Polygons. It can be used to check for
// collisions between objects without having to find out what type they are.
type Collider interface {
	OnPoint(*Point) bool
	OnBounding(*Bounding) bool
	OnLine(*Line) bool
	OnPolygon(*Polygon) bool
	DistanceToTangentPoint(*Point, Direction) (float64, float64)
}

// Collides checks if two Collider-satisfying objects are overlapping.
func Collides(collider1, collider2 Collider) bool {

	switch collider2.(type) {
	case *Point:
		return collider1.OnPoint(collider2.(*Point))
	case *Bounding:
		return collider1.OnBounding(collider2.(*Bounding))
	case *Line:
		return collider1.OnLine(collider2.(*Line))
	case *Polygon:
		return collider1.OnPolygon(collider2.(*Polygon))
	default:
		return false
	}
}
