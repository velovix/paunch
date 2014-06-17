package paunch

import (
	"math"
)

// Polygon is an object that represents a series of connected Lines that form
// a shape.
type Polygon struct {
	Lines  []*Line
	bounds *Bounding
}

// NewPolygon creates a new Polygon object.
func NewPolygon(points []*Point) *Polygon {

	var polygon Polygon
	polygon.Lines = make([]*Line, len(points))

	min := NewPoint(math.Inf(1), math.Inf(1))
	max := NewPoint(math.Inf(-1), math.Inf(-1))

	for i := 0; i < len(points); i++ {
		if i < len(points)-1 {
			polygon.Lines[i] = NewLine(points[i], points[i+1])
		} else {
			polygon.Lines[i] = NewLine(points[i], points[0])
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

	for _, val := range polygon.Lines {
		val.Move(x, y)
	}

	polygon.bounds.Move(x, y)
}

// SetPosition sets the position of the Polygon object with the first specified
// vertex as the start point.
func (polygon *Polygon) SetPosition(x, y float64) {

	xDisp := x - polygon.Lines[0].Start.X
	yDisp := y - polygon.Lines[0].Start.Y

	polygon.Move(xDisp, yDisp)
}

// GetPosition returns the starting position of the first line of the Polygon
// object.
func (polygon *Polygon) GetPosition() (x, y float64) {

	return polygon.Lines[0].Start.X, polygon.Lines[0].Start.Y
}

// DistanceToTangentPoint returns a Point with values equal to the distance
// a given Point is from the closest tangent Point on the given side of the
// Polygon.
func (polygon *Polygon) DistanceToTangentPoint(point *Point, side Direction) (float64, float64) {

	switch side {
	case Up:
		top := NewPoint(point.X, math.Inf(-1))
		for _, val := range polygon.Lines {
			linePnt, err := val.GetPointFromX(point.X)
			if gpfErr, ok := err.(LineGetPointFromError); linePnt.Y > top.Y && (!ok || gpfErr.Type != OutsideLineRangeError) {
				top = linePnt
			}
		}

		if !math.IsInf(top.Y, 0) {
			return getPointDistance(point, top)
		}

		for _, val := range polygon.Lines {
			linePnt, _ := val.GetPointFromX(point.X)
			if linePnt.Y > top.Y {
				top = linePnt
			}
		}
		return getPointDistance(point, top)
	case Down:
		bottom := NewPoint(point.X, math.Inf(1))
		for _, val := range polygon.Lines {
			linePnt, err := val.GetPointFromX(point.X)
			if gpfErr, ok := err.(LineGetPointFromError); linePnt.Y < bottom.Y && (!ok || gpfErr.Type != OutsideLineRangeError) {
				bottom = linePnt
			}
		}

		if !math.IsInf(bottom.Y, 0) {
			return getPointDistance(point, bottom)
		}

		for _, val := range polygon.Lines {
			linePnt, _ := val.GetPointFromX(point.X)
			if linePnt.Y < bottom.Y {
				bottom = linePnt
			}
		}
		return getPointDistance(point, bottom)
	case Left:
		left := NewPoint(math.Inf(1), point.Y)
		for _, val := range polygon.Lines {
			linePnt, err := val.GetPointFromY(point.Y)
			if gpfErr, ok := err.(LineGetPointFromError); linePnt.X < left.X && (!ok || gpfErr.Type != OutsideLineRangeError) {
				left = linePnt
			}
		}

		if !math.IsInf(left.X, 0) {
			return getPointDistance(point, left)
		}

		for _, val := range polygon.Lines {
			linePnt, _ := val.GetPointFromY(point.Y)
			if linePnt.X < left.X {
				left = linePnt
			}
		}
		return getPointDistance(point, left)
	case Right:
		right := NewPoint(math.Inf(-1), point.Y)
		for _, val := range polygon.Lines {
			linePnt, err := val.GetPointFromY(point.Y)
			if gpfErr, ok := err.(LineGetPointFromError); linePnt.X > right.X && (!ok || gpfErr.Type != OutsideLineRangeError) {
				right = linePnt
			}
		}

		if !math.IsInf(right.X, 0) {
			return getPointDistance(point, right)
		}

		for _, val := range polygon.Lines {
			linePnt, _ := val.GetPointFromY(point.Y)
			if linePnt.X > right.X {
				right = linePnt
			}
		}
		return getPointDistance(point, right)
	default:
		return 0, 0
	}
}

// OnPoint checks if a Point is on the Polygon object.
func (polygon *Polygon) OnPoint(point *Point) bool {

	if !point.OnBounding(polygon.bounds) {
		return false
	}

	ray := NewLine(NewPoint(math.Floor(point.X), math.Floor(point.Y)),
		NewPoint(math.Floor(point.X+(polygon.bounds.End.X-polygon.bounds.Start.X)), math.Floor(point.Y)))

	intersects := 0
	for _, val := range polygon.Lines {
		intersectPnt := getLineIntersection(ray, val)
		if intersectPnt == nil {
			continue
		} else {
			isOnVertex := false
			if intersectPnt.OnPoint(val.Start) && val.End.Y < intersectPnt.Y {
				intersects++
				isOnVertex = true
			} else if intersectPnt.OnPoint(val.End) && val.Start.Y < intersectPnt.Y {
				intersects++
				isOnVertex = true
			} else if intersectPnt.OnPoint(val.End) || intersectPnt.OnPoint(val.Start) {
				isOnVertex = true
			}

			if !isOnVertex {
				intersects++
			}
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

	for _, val := range polygon.Lines {
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

	for _, val := range polygon.Lines {
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
