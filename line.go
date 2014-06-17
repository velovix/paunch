package paunch

import (
	"math"
)

// Line is an object that represents a line segment.
type Line struct {
	Start  *Point
	End    *Point
	bounds *Bounding

	M float64
	B float64
}

// NewLine creates a new line object. This is absolutely necissary before use.
func NewLine(start, end *Point) *Line {

	var calcStart, calcEnd *Point
	if start.X > end.X {
		calcStart = NewPoint(end.X, end.Y)
		calcEnd = NewPoint(start.X, start.Y)
	} else {
		calcStart = NewPoint(start.X, start.Y)
		calcEnd = NewPoint(end.X, end.Y)
	}

	line := &Line{Start: calcStart, End: calcEnd}

	line.bounds = NewBounding(line.Start, line.End)

	line.M = getSlope(line.Start.X, line.Start.Y, line.End.X, line.End.Y)
	line.B = line.Start.Y - (line.M * line.Start.X)

	return line
}

// Move moves the Line object a specified distance.
func (line *Line) Move(x, y float64) {

	line.Start.Move(x, y)
	line.End.Move(x, y)

	line.bounds.Move(x, y)

	line.B = line.Start.Y - (line.M * line.Start.X)
}

// SetPosition sets the position of the Line object with the start point as the
// reference point.
func (line *Line) SetPosition(x, y float64) {

	xDisp := x - line.Start.X
	yDisp := y - line.Start.Y

	line.Move(xDisp, yDisp)
}

// GetPosition returns the starting position of the Line object.
func (line *Line) GetPosition() (x, y float64) {

	return line.Start.X, line.Start.Y
}

// GetPointFromX returns a Point on the Line that corresponds to the given X
// value. If the given X value is outside the area of the line, the method will
// return the nearest Point and an error. If the slope of the line is
// undefined, the method will return the highest Point on the Line and an
// error.
func (line *Line) GetPointFromX(x float64) (*Point, error) {

	if x < line.bounds.Start.X {
		return NewPoint(line.Start.X, line.Start.Y), LineGetPointFromError{x, line, OutsideLineRangeError}
	} else if x > line.bounds.End.X {
		return NewPoint(line.End.X, line.End.Y), LineGetPointFromError{x, line, OutsideLineRangeError}
	}

	if math.IsInf(line.M, 0) {
		return NewPoint(line.End.X, line.End.Y), LineGetPointFromError{x, line, UndefinedSlopeError}
	}

	return NewPoint(x, (line.M*x)+line.B), nil
}

// GetPointFromY returns a Point on the Line that corresponds to the given Y
// value. If the given Y value is outside the area of the line, the method will
// return the nearest Point and an error.
func (line *Line) GetPointFromY(y float64) (*Point, error) {

	if y < line.bounds.Start.Y || y > line.bounds.End.Y {
		if math.Abs(line.Start.Y-y) < math.Abs(line.End.Y-y) {
			return NewPoint(line.Start.X, line.Start.Y), LineGetPointFromError{y, line, OutsideLineRangeError}
		}

		return NewPoint(line.End.X, line.End.Y), LineGetPointFromError{y, line, OutsideLineRangeError}
	}

	if math.IsInf(line.M, 0) {
		return NewPoint(line.Start.X, y), nil
	}

	return NewPoint((y-line.B)/line.M, y), nil
}

// DistanceToTangentPoint returns a Point with values equal to the distance
// a given Point is from the closest tangent Point on the given side of the
// Line.
func (line *Line) DistanceToTangentPoint(point *Point, side Direction) (float64, float64) {

	switch side {
	case Up:
		x := point.X
		tangent, _ := line.GetPointFromX(x)
		return getPointDistance(point, tangent)
	case Down:
		x := point.X
		tangent, _ := line.GetPointFromX(x)
		return getPointDistance(point, tangent)
	case Left:
		y := point.Y
		tangent, _ := line.GetPointFromY(y)
		return getPointDistance(point, tangent)
	case Right:
		y := point.Y
		tangent, _ := line.GetPointFromY(y)
		return getPointDistance(point, tangent)
	default:
		return 0, 0
	}
}

func (bounding *Bounding) getLines() []*Line {

	line := make([]*Line, 4)

	line[0] = NewLine(NewPoint(bounding.Start.X, bounding.Start.Y), NewPoint(bounding.End.X, bounding.Start.Y))
	line[1] = NewLine(NewPoint(bounding.End.X, bounding.Start.Y), NewPoint(bounding.End.X, bounding.End.Y))
	line[2] = NewLine(NewPoint(bounding.End.X, bounding.End.Y), NewPoint(bounding.Start.X, bounding.End.Y))
	line[3] = NewLine(NewPoint(bounding.Start.X, bounding.End.Y), NewPoint(bounding.Start.X, bounding.Start.Y))

	return line
}

// OnPoint checks if a Point is on the Line object.
func (line *Line) OnPoint(point *Point) bool {

	if math.IsInf(line.M, 0) {
		if point.Y >= line.bounds.Start.Y && point.Y <= line.bounds.End.Y &&
			math.Abs(point.X-line.Start.X) < tolerance {
			return true
		}

		return false
	}

	if !point.OnBounding(line.bounds) {
		return false
	}

	if math.Abs(point.Y-((line.M*point.X)+line.B)) < tolerance {
		return true
	}

	return false
}

// OnBounding checks if a Bounding is on the Line object.
func (line *Line) OnBounding(bounding *Bounding) bool {

	if !bounding.OnBounding(line.bounds) {
		return false
	}

	if line.Start.OnBounding(bounding) || line.End.OnBounding(bounding) {
		return true
	}

	boundLines := bounding.getLines()
	for _, val := range boundLines {
		if line.OnLine(val) {
			return true
		}
	}

	return false
}

func getLineIntersection(line1, line2 *Line) *Point {

	if !line1.bounds.OnBounding(line2.bounds) {
		return nil
	}

	if line1.M == line2.M {
		return nil
	}

	line1A := line1.End.Y - line1.Start.Y
	line1B := line1.Start.X - line1.End.X
	line1C := line1A*line1.Start.X + line1B*line1.Start.Y

	line2A := line2.End.Y - line2.Start.Y
	line2B := line2.Start.X - line2.End.X
	line2C := line2A*line2.Start.X + line2B*line2.Start.Y

	determinate := line1A*line2B - line2A*line1B

	x := (line2B*line1C - line1B*line2C) / determinate
	y := (line1A*line2C - line2A*line1C) / determinate

	if x >= line1.Start.X && x <= line1.End.X && y >= line1.Start.Y && y <= line1.End.Y {
		return NewPoint(x, y)
	}

	return nil
}

// OnLine checks if a line is on the Line object.
func (line *Line) OnLine(line2 *Line) bool {

	if !line.bounds.OnBounding(line2.bounds) {
		return false
	}

	if intersection := getLineIntersection(line, line2); intersection != nil {
		return true
	}

	return false
}

// OnLine checks if a Line is on the Bounding object.
func (bounding *Bounding) OnLine(line *Line) bool {

	return line.OnBounding(bounding)
}

// OnLine checks if a Line is on the Point object.
func (point *Point) OnLine(line *Line) bool {

	return line.OnPoint(point)
}
