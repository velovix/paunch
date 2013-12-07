package paunch

// Collision is an interface that qualifies types that can check for collisions
// against Points, Boundings, Lines, and Polygons. It can be used to check for
// collisions between objects without having to find out what type they are.
type Collision interface {
	OnPoint(Point) bool
	OnBounding(Bounding) bool
	OnLine(Line) bool
	OnPolygon(Polygon) bool
	Move(x, y float64)
}

// NewCollision creates a new collision object who's type is determined
// automatically based on how many Points are given. (1: Point, 2: Line)
func NewCollision(points []Point, objType int) Collision {

	switch objType {
	case POINT:
		return Collision(points[0])
	case BOUNDING:
		return Collision(NewBounding(points[0], points[1]))
	case LINE:
		return Collision(NewLine(points[0], points[1]))
	case POLYGON:
		return Collision(NewPolygon(points))
	default:
		return nil
	}
}

func Collides(collision1, collision2 Collision) bool {

	switch collision2.(type) {
	case Point:
		return collision1.OnPoint(collision2.(Point))
	case Bounding:
		return collision1.OnBounding(collision2.(Bounding))
	case Line:
		return collision1.OnLine(collision2.(Line))
	case Polygon:
		return collision1.OnPolygon(collision2.(Polygon))
	default:
		return false
	}
}
