package paunch

// bounding is an object that represents a bounding box. It is meant to be
// used through the Collider interface.
type bounding struct {
	start *point
	end   *point
}

func newBounding(start *point, end *point) *bounding {

	var checkStart, checkEnd point
	if start.x >= end.x {
		checkEnd.x = start.x
		checkStart.x = end.x
	} else {
		checkEnd.x = end.x
		checkStart.x = start.x
	}
	if start.y >= end.y {
		checkEnd.y = start.y
		checkStart.y = end.y
	} else {
		checkEnd.y = end.y
		checkStart.y = start.y
	}

	return &bounding{start: &checkStart, end: &checkEnd}
}

func (b *bounding) Move(x, y float64) {

	b.start.Move(x, y)
	b.end.Move(x, y)
}

func (b *bounding) SetPosition(x, y float64) {

	width := b.end.x - b.start.x
	height := b.end.y - b.start.y

	newBounding := newBounding(newPoint(x, y), newPoint(x+width, y+height))
	*b = *newBounding
}

func (b *bounding) GetPosition() (x, y float64) {

	return b.start.x, b.start.y
}

func (b *bounding) DistanceToTangentPoint(x, y float64, side Direction) (float64, float64) {

	switch side {
	case Up:
		sideX := x
		if x < b.start.x {
			sideX = b.start.x
		} else if x > b.end.x {
			sideX = b.end.x
		}
		return getPointDistance(newPoint(x, y), newPoint(sideX, b.end.y))
	case Down:
		sideX := x
		if x < b.start.x {
			sideX = b.start.x
		} else if x > b.end.x {
			sideX = b.end.x
		}
		return getPointDistance(newPoint(x, y), newPoint(sideX, b.start.y))
	case Left:
		sideY := y
		if y < b.start.y {
			sideY = b.start.y
		} else if y > b.end.y {
			sideY = b.end.y
		}
		return getPointDistance(newPoint(x, y), newPoint(b.start.x, sideY))
	case Right:
		sideY := y
		if y < b.start.y {
			sideY = b.start.y
		} else if y > b.end.y {
			sideY = b.end.y
		}
		return getPointDistance(newPoint(x, y), newPoint(b.end.x, sideY))
	default:
		return 0, 0
	}
}

func (b *bounding) onPoint(p *point) bool {

	if p.x >= b.start.x && p.x <= b.end.x &&
		p.y >= b.start.y && p.y <= b.end.y {
		return true
	}

	return false
}

func (b *bounding) onBounding(b2 *bounding) bool {

	if b.start.x > b2.end.x || b.end.x < b2.start.x ||
		b.start.y > b2.end.y || b.end.y < b2.start.y {
		return false
	}

	return true
}

func (p *point) onBounding(b *bounding) bool {

	return b.onPoint(p)
}
