package paunch

import (
	"math"
)

// Polygon is an object that represents a series of connected Lines that form
// a shape.
type Polygon struct {
	lines  []*Line
	bounds *Bounding
}

// NewPolygon creates a new Polygon object.
func NewPolygon(points []*Point) *Polygon {

	var polygon Polygon
	polygon.lines = make([]*Line, len(points))

	min := NewPoint(math.Inf(1), math.Inf(1))
	max := NewPoint(math.Inf(-1), math.Inf(-1))

	for i := 0; i < len(points); i++ {
		if i < len(points)-1 {
			polygon.lines[i] = NewLine(points[i], points[i+1])
		} else {
			polygon.lines[i] = NewLine(points[i], points[0])
		}

		if points[i].X > max.X {
			max.X = points[i].X
		}
		if points[i].Y > max.Y {
			max.Y = points[i].Y
		}
		if points[i].X < min.X {
			min.X = points[i].X
		}
		if points[i].Y < min.Y {
			min.Y = points[i].Y
		}
	}

	polygon.bounds = NewBounding(min, max)
	return &polygon
}

// Move moves the Polygon object a specified distance.
func (polygon *Polygon) Move(x, y float64) {

	for i := range polygon.lines {
		polygon.lines[i].Move(x, y)
	}

	polygon.bounds.Move(x, y)
}

// DistanceToTangentPoint returns a Point with values equal to the distance
// a given Point is from the closest tangent Point on the given side of the
// Polygon.
func (polygon *Polygon) DistanceToTangentPoint(point *Point, side Direction) *Point {

	switch side {
	case Up:
		top := NewPoint(point.X, math.Inf(-1))
		for _, val := range polygon.lines {
			linePnt, err := val.GetPointFromX(point.X)
			if gpfErr, ok := err.(LineGetPointFromError); linePnt.Y > top.Y && (!ok || gpfErr.Type != OutsideLineRangeError) {
				top = linePnt
			}
		}
		return getPointDistance(point, top)
	case Down:
		bottom := NewPoint(point.X, math.Inf(1))
		for _, val := range polygon.lines {
			linePnt, err := val.GetPointFromX(point.X)
			if gpfErr, ok := err.(LineGetPointFromError); linePnt.Y < bottom.Y && (!ok || gpfErr.Type != OutsideLineRangeError) {
				bottom = linePnt
			}
		}
		return getPointDistance(point, bottom)
	case Left:
		left := NewPoint(math.Inf(1), point.Y)
		for _, val := range polygon.lines {
			linePnt, err := val.GetPointFromY(point.Y)
			if gpfErr, ok := err.(LineGetPointFromError); linePnt.X < left.X && (!ok || gpfErr.Type != OutsideLineRangeError) {
				left = linePnt
			}
		}
		return getPointDistance(point, left)
	case Right:
		right := NewPoint(math.Inf(-1), point.Y)
		for _, val := range polygon.lines {
			linePnt, err := val.GetPointFromY(point.Y)
			if gpfErr, ok := err.(LineGetPointFromError); linePnt.X > right.X && (!ok || gpfErr.Type != OutsideLineRangeError) {
				right = linePnt
			}
		}
		return getPointDistance(point, right)
	default:
		return NewPoint(0, 0)
	}
}

// OnPoint checks if a Point is on the Polygon object.
func (polygon *Polygon) OnPoint(point *Point) bool {

	if !point.OnBounding(polygon.bounds) {
		return false
	}

	ray := NewLine(point, NewPoint(point.X+999, point.Y))

	intersects := 0
	for _, val := range polygon.lines {
		if ray.OnLine(val) {
			intersects++
		}
	}

	if intersects%2 == 0 {
		return false
	}

	return true
}

// OnBounding checks if a Bounding is on the Polygon object.
func (polygon *Polygon) OnBounding(bounding *Bounding) bool {

	if !bounding.OnBounding(polygon.bounds) {
		return false
	}

	if polygon.OnPoint(bounding.Start) || polygon.OnPoint(bounding.End) ||
		polygon.OnPoint(NewPoint(bounding.Start.X, bounding.End.Y)) ||
		polygon.OnPoint(NewPoint(bounding.End.X, bounding.Start.Y)) {
		return true
	}

	boundLines := bounding.getLines()
	for _, val := range boundLines {
		if polygon.OnLine(val) {
			return true
		}
	}

	return false
}

// OnLine checks if a Line is on the Polygon object.
func (polygon *Polygon) OnLine(line *Line) bool {

	if !line.bounds.OnBounding(polygon.bounds) {
		return false
	}

	for _, val := range polygon.lines {
		if line.OnLine(val) {
			return true
		}
	}

	return false
}

// OnPolygon checks if a Polygon is on the Polygon object.
func (polygon *Polygon) OnPolygon(polygon2 *Polygon) bool {

	if !polygon.bounds.OnBounding(polygon2.bounds) {
		return false
	}

	for _, val := range polygon.lines {
		if polygon2.OnLine(val) {
			return true
		}
	}

	return false
}

// OnPolygon checks if a Polygon is on the Line object.
func (line *Line) OnPolygon(polygon *Polygon) bool {

	return polygon.OnLine(line)
}

// OnPolygon checks if a Polygon is on the Bounding object.
func (bounding *Bounding) OnPolygon(polygon *Polygon) bool {

	return polygon.OnBounding(bounding)
}

// OnPolygon checks if a Polygon is on the Point object.
func (point *Point) OnPolygon(polygon *Polygon) bool {

	return polygon.OnPoint(point)
}
