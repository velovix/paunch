package paunch

import (
	"math"
)

// line is an object that represents a line segment. It is meant to be used
// through the collider interface.
type line struct {
	start  *point
	end    *point
	bounds *bounding

	m float64
	b float64
}

func newLine(start, end *point) *line {

	var calcStart, calcEnd *point
	if start.x > end.x {
		calcStart = newPoint(end.x, end.y)
		calcEnd = newPoint(start.x, start.y)
	} else {
		calcStart = newPoint(start.x, start.y)
		calcEnd = newPoint(end.x, end.y)
	}

	l := &line{start: calcStart, end: calcEnd}

	l.bounds = newBounding(l.start, l.end)

	l.m = getSlope(l.start.x, l.start.y, l.end.x, l.end.y)
	l.b = l.start.y - (l.m * l.start.x)

	return l
}

func (l *line) Move(x, y float64) {

	l.start.Move(x, y)
	l.end.Move(x, y)

	l.bounds.Move(x, y)

	l.b = l.start.y - (l.m * l.start.x)
}

func (l *line) SetPosition(x, y float64) {

	xDisp := x - l.start.x
	yDisp := y - l.start.y

	l.Move(xDisp, yDisp)
}

func (l *line) Position() (x, y float64) {

	return l.start.x, l.start.y
}

func (l *line) getPointFromX(x float64) (*point, error) {

	if x < l.bounds.start.x {
		return newPoint(l.start.x, l.start.y), lineGetPointFromError{x, l, outsideLineRangeError}
	} else if x > l.bounds.end.x {
		return newPoint(l.end.x, l.end.y), lineGetPointFromError{x, l, outsideLineRangeError}
	}

	if math.IsInf(l.m, 0) {
		return newPoint(l.end.x, l.end.y), lineGetPointFromError{x, l, undefinedSlopeError}
	}

	return newPoint(x, (l.m*x)+l.b), nil
}

func (l *line) getPointFromY(y float64) (*point, error) {

	if y < l.bounds.start.y || y > l.bounds.end.y {
		if math.Abs(l.start.y-y) < math.Abs(l.end.y-y) {
			return newPoint(l.start.x, l.start.y), lineGetPointFromError{y, l, outsideLineRangeError}
		}

		return newPoint(l.end.x, l.end.y), lineGetPointFromError{y, l, outsideLineRangeError}
	}

	if math.IsInf(l.m, 0) {
		return newPoint(l.start.x, y), nil
	}

	return newPoint((y-l.b)/l.m, y), nil
}

func (l *line) DistanceToTangentPoint(x, y float64, side Direction) (float64, float64) {

	switch side {
	case Up:
		tangent, _ := l.getPointFromX(x)
		return getPointDistance(newPoint(x, y), tangent)
	case Down:
		tangent, _ := l.getPointFromX(x)
		return getPointDistance(newPoint(x, y), tangent)
	case Left:
		tangent, _ := l.getPointFromY(y)
		return getPointDistance(newPoint(x, y), tangent)
	case Right:
		tangent, _ := l.getPointFromY(y)
		return getPointDistance(newPoint(x, y), tangent)
	default:
		return 0, 0
	}
}

func (b *bounding) getLines() []*line {

	line := make([]*line, 4)

	line[0] = newLine(newPoint(b.start.x, b.start.y), newPoint(b.end.x, b.start.y))
	line[1] = newLine(newPoint(b.end.x, b.start.y), newPoint(b.end.x, b.end.y))
	line[2] = newLine(newPoint(b.end.x, b.end.y), newPoint(b.start.x, b.end.y))
	line[3] = newLine(newPoint(b.start.x, b.end.y), newPoint(b.start.x, b.start.y))

	return line
}

func (l *line) onPoint(p *point) bool {

	if math.IsInf(l.m, 0) {
		if p.y >= l.bounds.start.y && p.y <= l.bounds.end.y &&
			math.Abs(p.x-l.start.x) < tolerance {
			return true
		}

		return false
	}

	if !p.onBounding(l.bounds) {
		return false
	}

	if math.Abs(p.y-((l.m*p.x)+l.b)) < tolerance {
		return true
	}

	return false
}

func (l *line) onBounding(b *bounding) bool {

	if !b.onBounding(l.bounds) {
		return false
	}

	if l.start.onBounding(b) || l.end.onBounding(b) {
		return true
	}

	boundlines := b.getLines()
	for _, val := range boundlines {
		if l.onLine(val) {
			return true
		}
	}

	return false
}

func getLineIntersection(l1, l2 *line) *point {

	if !l1.bounds.onBounding(l2.bounds) {
		return nil
	}

	if l1.m == l2.m {
		return nil
	}

	l1A := l1.end.y - l1.start.y
	l1B := l1.start.x - l1.end.x
	l1C := l1A*l1.start.x + l1B*l1.start.y

	l2A := l2.end.y - l2.start.y
	l2B := l2.start.x - l2.end.x
	l2C := l2A*l2.start.x + l2B*l2.start.y

	determinate := l1A*l2B - l2A*l1B

	x := (l2B*l1C - l1B*l2C) / determinate
	y := (l1A*l2C - l2A*l1C) / determinate

	if x >= l1.start.x && x <= l1.end.x && y >= l1.start.y && y <= l1.end.y {
		return newPoint(x, y)
	}

	return nil
}

func (l *line) onLine(l2 *line) bool {

	if !l.bounds.onBounding(l2.bounds) {
		return false
	}

	if intersection := getLineIntersection(l, l2); intersection != nil {
		return true
	}

	return false
}

func (b *bounding) onLine(l *line) bool {

	return l.onBounding(b)
}

func (p *point) onLine(l *line) bool {

	return l.onPoint(p)
}
