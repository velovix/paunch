package paunch

// point is an object that represents an X and Y position in 2D space. It is
// meant to be used through the Collider interface.
type point struct {
	x float64
	y float64
}

func getPointDistance(point1, point2 *point) (float64, float64) {

	x, y := point2.x-point1.x, point2.y-point1.y

	return x, y
}

func newPoint(x, y float64) *point {

	return &point{x: x, y: y}
}

func (p *point) Move(x, y float64) {

	p.x += x
	p.y += y
}

func (p *point) SetPosition(x, y float64) {

	xDisp := x - p.x
	yDisp := y - p.y

	p.Move(xDisp, yDisp)
}

func (p *point) Position() (x, y float64) {

	return p.x, p.y
}

func (p *point) DistanceToTangentPoint(x, y float64, side Direction) (float64, float64) {

	return getPointDistance(newPoint(x, y), p)
}

func (p *point) onPoint(p2 *point) bool {

	if p.x == p2.x && p.y == p2.y {
		return true
	}

	return false
}
