package paunch

import (
	"math"
)

type Line struct {
	Start  Point
	End    Point
	bounds Bounding

	M float64
	B float64
}

func NewLine(start, end Point) Line {

	line := Line{Start: start, End: end}

	line.bounds = NewBounding(line.Start, line.End)

	line.M = GetSlope(line.Start, line.End)
	line.B = line.Start.Y - (line.M * line.Start.X)

	return line
}

func (bounding Bounding) getLines() []Line {

	line := make([]Line, 4)

	line[0] = NewLine(NewPoint(bounding.Start.X, bounding.Start.Y), NewPoint(bounding.End.X, bounding.Start.Y))
	line[1] = NewLine(NewPoint(bounding.End.X, bounding.Start.Y), NewPoint(bounding.End.X, bounding.End.Y))
	line[2] = NewLine(NewPoint(bounding.End.X, bounding.End.Y), NewPoint(bounding.Start.X, bounding.End.Y))
	line[3] = NewLine(NewPoint(bounding.Start.X, bounding.End.Y), NewPoint(bounding.Start.X, bounding.Start.Y))

	return line
}

func isPointOnLine(point Point, line Line) bool {

	if math.IsInf(line.M, 0) {
		if point.Y >= line.bounds.Start.Y && point.Y <= line.bounds.End.Y &&
			math.Abs(point.X-line.Start.X) < TOLERANCE {
			return true
		} else {
			return false
		}
	}

	if !isPointOnBounding(point, line.bounds) {
		return false
	}

	if math.Abs(point.Y-((line.M*point.X)+line.B)) < TOLERANCE {
		return true
	}

	return false
}

func isBoundingOnLine(bounding Bounding, line Line) bool {

	if !isBoundingOnBounding(bounding, line.bounds) {
		return false
	}

	if isPointOnBounding(line.Start, bounding) || isPointOnBounding(line.End, bounding) {
		return true
	}

	boundLines := bounding.getLines()
	for _, val := range boundLines {
		if isLineOnLine(line, val) {
			return true
		}
	}

	return false
}

func isLineOnLine(line1, line2 Line) bool {

	if !isBoundingOnBounding(line1.bounds, line2.bounds) {
		return false
	}

	dx1 := FindDeterminate(
		FindDeterminate(line1.Start.X, line1.Start.Y, line1.End.X, line1.End.Y),
		FindDeterminate(line1.Start.X, 1, line1.End.X, 1),
		FindDeterminate(line2.Start.X, line2.Start.Y, line2.End.X, line2.End.Y),
		FindDeterminate(line2.Start.X, 1, line2.End.X, 1))

	dxy2 := FindDeterminate(
		FindDeterminate(line1.Start.X, 1, line1.End.X, 1),
		FindDeterminate(line1.Start.Y, 1, line1.End.Y, 1),
		FindDeterminate(line2.Start.X, 1, line2.End.X, 1),
		FindDeterminate(line2.Start.Y, 1, line2.End.Y, 1))

	dy1 := FindDeterminate(
		FindDeterminate(line1.Start.X, line1.Start.Y, line1.End.X, line1.End.Y),
		FindDeterminate(line1.Start.Y, 1, line1.End.Y, 1),
		FindDeterminate(line2.Start.X, line2.Start.Y, line2.End.X, line2.End.Y),
		FindDeterminate(line2.Start.Y, 1, line2.End.Y, 1))

	x := dx1 / dxy2
	y := dy1 / dxy2

	if isPointOnLine(NewPoint(x, y), line1) && isPointOnLine(NewPoint(x, y), line2) {
		return true
	}

	return false
}
