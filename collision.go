package paunch

// Collision is an interface that qualifies types that can check for collisions
// against Points, Boundings, Lines, and Polygons. It can be used to check for
// collisions between objects without having to find out what type they are.
type Collision interface {
	OnPoint(*Point) bool
	OnBounding(*Bounding) bool
	OnLine(*Line) bool
	OnPolygon(*Polygon) bool
	Move(x, y float64)
}

// Collides checks if two Collision objects are overlapping.
func Collides(collision1, collision2 Collision) bool {

	switch collision2.(type) {
	case *Point:
		return collision1.OnPoint(collision2.(*Point))
	case *Bounding:
		return collision1.OnBounding(collision2.(*Bounding))
	case *Line:
		return collision1.OnLine(collision2.(*Line))
	case *Polygon:
		return collision1.OnPolygon(collision2.(*Polygon))
	default:
		return false
	}
}
