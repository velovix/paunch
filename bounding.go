package paunch

// Bounding is an object that represents a bounding box.
type Bounding struct {
	Start *Point
	End   *Point
}

// NewBounding creates a new Bounding object.
func NewBounding(start *Point, end *Point) *Bounding {

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

	return &Bounding{Start: &checkStart, End: &checkEnd}
}

// Move moves the Bounding object a specified distance.
func (bounding *Bounding) Move(x, y float64) {

	bounding.Start.Move(x, y)
	bounding.End.Move(x, y)
}

// SetPosition sets the position of the Bounding object with the start point as
// the reference point.
func (bounding *Bounding) SetPosition(x, y float64) {

	width := bounding.End.X - bounding.Start.X
	height := bounding.End.Y - bounding.Start.Y

	newBounding := NewBounding(NewPoint(x, y), NewPoint(x+width, y+height))
	*bounding = *newBounding
}

// GetPosition returns the bottom left position of the Bounding object.
func (bounding *Bounding) GetPosition() (x, y float64) {

	return bounding.Start.X, bounding.Start.Y
}

// DistanceToTangentPoint returns a Point with values equal to the distance
// a given Point is from the closest tangent Point on the given side of the
// Bounding.
func (bounding *Bounding) DistanceToTangentPoint(point *Point, side Direction) *Point {

	switch side {
	case Up:
		x := point.X
		if point.X < bounding.Start.X {
			x = bounding.Start.X
		} else if point.X > bounding.End.X {
			x = bounding.End.X
		}
		return getPointDistance(point, NewPoint(x, bounding.End.Y))
	case Down:
		x := point.X
		if point.X < bounding.Start.X {
			x = bounding.Start.X
		} else if point.X > bounding.End.X {
			x = bounding.End.X
		}
		return getPointDistance(point, NewPoint(x, bounding.Start.Y))
	case Left:
		y := point.Y
		if point.Y < bounding.Start.Y {
			y = bounding.Start.Y
		} else if point.Y > bounding.End.Y {
			y = bounding.End.Y
		}
		return getPointDistance(point, NewPoint(bounding.Start.X, y))
	case Right:
		y := point.Y
		if point.Y < bounding.Start.Y {
			y = bounding.Start.Y
		} else if point.Y > bounding.End.Y {
			y = bounding.End.Y
		}
		return getPointDistance(point, NewPoint(bounding.End.X, y))
	default:
		return NewPoint(0, 0)
	}
}

// OnPoint checks if a Point is on the Bounding object.
func (bounding *Bounding) OnPoint(point *Point) bool {

	if point.X >= bounding.Start.X && point.X <= bounding.End.X &&
		point.Y >= bounding.Start.Y && point.Y <= bounding.End.Y {
		return true
	}

	return false
}

// OnBounding checks if a Bounding is on the Bounding object.
func (bounding *Bounding) OnBounding(bounding2 *Bounding) bool {

	if bounding.Start.X > bounding2.End.X || bounding.End.X < bounding2.Start.X ||
		bounding.Start.Y > bounding2.End.Y || bounding.End.Y < bounding2.Start.Y {
		return false
	}

	return true
}

// OnBounding checks if a Bounding is on the Point object.
func (point *Point) OnBounding(bounding *Bounding) bool {

	return bounding.OnPoint(point)
}
