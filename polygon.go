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

	for i, _ := range polygon.lines {
		polygon.lines[i].Move(x, y)
	}

	polygon.bounds.Move(x, y)
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
	} else {
		return true
	}
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
func (polygon1 *Polygon) OnPolygon(polygon2 *Polygon) bool {

	if !polygon1.bounds.OnBounding(polygon2.bounds) {
		return false
	}

	for _, val := range polygon1.lines {
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
