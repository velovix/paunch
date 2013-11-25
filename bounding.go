package paunch

type Bounding struct {
	Start Point
	End   Point
}

func NewBounding(start Point, end Point) Bounding {

	var checkStart, checkEnd Point
	if start.X >= end.X {
		checkEnd.X = start.X
		checkStart.X = end.X
	} else {
		checkEnd.X = end.X
		checkStart.X = start.X
	}
	if start.Y >= end.Y {
		checkEnd.Y = start.Y
		checkStart.Y = end.Y
	} else {
		checkEnd.Y = end.Y
		checkStart.Y = start.Y
	}

	return Bounding{Start: checkStart, End: checkEnd}
}

func isPointOnBounding(point Point, bounding Bounding) bool {

	if point.X >= bounding.Start.X && point.X <= bounding.End.X &&
		point.Y >= bounding.Start.Y && point.Y <= bounding.End.Y {
		return true
	}

	return false
}

func isBoundingOnBounding(bounding1, bounding2 Bounding) bool {

	if bounding1.Start.X > bounding2.End.X || bounding1.End.X < bounding2.Start.X ||
		bounding1.Start.Y > bounding1.End.Y || bounding1.End.Y < bounding2.Start.Y {
		return false
	}

	return true
}
