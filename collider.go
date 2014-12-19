package paunch

import (
	"math"
)

// Collider is an object that represents a shape that can be tested for
// collision against another Collider.
type Collider interface {
	onPoint(*point) bool
	onBounding(*bounding) bool
	onLine(*line) bool
	onPolygon(*polygon) bool

	// Move moves the Collider object the specified distance.
	Move(x, y float64)
	// SetPosition sets the position of the Collider.
	SetPosition(x, y float64)
	// Position returns the x, y coordinates of the Collider object's
	// current position.
	Position() (float64, float64)
	// DistanceToTangentPoint returns the x, y coordinates of the nearest point
	// tangent to the Collider object. This method is useful for position
	// correction when objects have sunk into each other.
	DistanceToTangentPoint(float64, float64, Direction) (float64, float64)
}

// NewCollider creates a new Collider object. The supplied coordinates should
// be in an "x1, y1, x2, y2..." format. Colliders work differently internally
// depending on the shape the coordinate describes. Collision detection is
// faster for singular points and bounding boxes than with lines and polygons.
func NewCollider(coords []float64) Collider {

	if len(coords) == 0 || len(coords)%2 != 0 {
		return nil
	}

	if len(coords) == 2 {
		return newPoint(coords[0], coords[1])
	}

	if len(coords) == 4 {
		return newLine(newPoint(coords[0], coords[1]), newPoint(coords[2], coords[3]))
	}

	if len(coords) == 8 {
		// Check to see that the coordinates have two unique x and y values
		var uniqueX, uniqueY []float64
		for i, val := range coords {
			taken := false
			if i%2 == 0 {
				for _, val2 := range uniqueX {
					if val == val2 {
						taken = true
					}
				}
				if !taken {
					uniqueX = append(uniqueX, val)
				}
			} else {
				for _, val2 := range uniqueY {
					if val == val2 {
						taken = true
					}
				}
				if !taken {
					uniqueY = append(uniqueY, val)
				}
			}
		}

		if len(uniqueX) == 2 && len(uniqueY) == 2 {
			return newBounding(newPoint(math.Min(uniqueX[0], uniqueX[1]), math.Min(uniqueY[0], uniqueY[1])),
				newPoint(math.Max(uniqueX[0], uniqueX[1]), math.Max(uniqueY[0], uniqueY[1])))
		}
	}

	points := make([]*point, len(coords)/2)
	for i := 0; i < len(coords); i += 2 {
		points[i/2] = newPoint(coords[i], coords[i+1])
	}

	return newPolygon(points)
}

// Collides checks if two Collider-satisfying objects are overlapping.
func Collides(collider1, collider2 Collider) bool {

	switch collider2.(type) {
	case *point:
		return collider1.onPoint(collider2.(*point))
	case *bounding:
		return collider1.onBounding(collider2.(*bounding))
	case *line:
		return collider1.onLine(collider2.(*line))
	case *polygon:
		return collider1.onPolygon(collider2.(*polygon))
	default:
		return false
	}
}
