package paunch

import (
	"math"
)

// polygon is an object that represents a series of connected lines that form
// a shape. It is meant to be used through the Collider interface.
type polygon struct {
	lines  []*line
	bounds *bounding
}

func newPolygon(points []*point) *polygon {

	var poly polygon
	poly.lines = make([]*line, len(points))

	min := newPoint(math.Inf(1), math.Inf(1))
	max := newPoint(math.Inf(-1), math.Inf(-1))

	for i := 0; i < len(points); i++ {
		if i < len(points)-1 {
			poly.lines[i] = newLine(points[i], points[i+1])
		} else {
			poly.lines[i] = newLine(points[i], points[0])
		}

		if points[i].x > max.x {
			max.x = points[i].x
		}
		if points[i].y > max.y {
			max.y = points[i].y
		}
		if points[i].x < min.x {
			min.x = points[i].x
		}
		if points[i].y < min.y {
			min.y = points[i].y
		}
	}

	poly.bounds = newBounding(min, max)
	return &poly
}

func (poly *polygon) Move(x, y float64) {

	for _, val := range poly.lines {
		val.Move(x, y)
	}

	poly.bounds.Move(x, y)
}

func (poly *polygon) SetPosition(x, y float64) {

	xDisp := x - poly.lines[0].start.x
	yDisp := y - poly.lines[0].start.y

	poly.Move(xDisp, yDisp)
}

func (poly *polygon) Position() (x, y float64) {

	return poly.lines[0].start.x, poly.lines[0].start.y
}

func (poly *polygon) DistanceToTangentPoint(x, y float64, side Direction) (float64, float64) {

	switch side {
	case Up:
		top := newPoint(x, math.Inf(-1))
		for _, val := range poly.lines {
			linePnt, err := val.getPointFromX(x)
			if gpfErr, ok := err.(lineGetPointFromError); linePnt.y > top.y && (!ok || gpfErr.Type != outsideLineRangeError) {
				top = linePnt
			}
		}

		if !math.IsInf(top.y, 0) {
			return getPointDistance(newPoint(x, y), top)
		}

		for _, val := range poly.lines {
			linePnt, _ := val.getPointFromX(x)
			if linePnt.y > top.y {
				top = linePnt
			}
		}
		return getPointDistance(newPoint(x, y), top)
	case Down:
		bottom := newPoint(x, math.Inf(1))
		for _, val := range poly.lines {
			linePnt, err := val.getPointFromX(x)
			if gpfErr, ok := err.(lineGetPointFromError); linePnt.y < bottom.y && (!ok || gpfErr.Type != outsideLineRangeError) {
				bottom = linePnt
			}
		}

		if !math.IsInf(bottom.y, 0) {
			return getPointDistance(newPoint(x, y), bottom)
		}

		for _, val := range poly.lines {
			linePnt, _ := val.getPointFromX(x)
			if linePnt.y < bottom.y {
				bottom = linePnt
			}
		}
		return getPointDistance(newPoint(x, y), bottom)
	case Left:
		left := newPoint(math.Inf(1), y)
		for _, val := range poly.lines {
			linePnt, err := val.getPointFromY(y)
			if gpfErr, ok := err.(lineGetPointFromError); linePnt.x < left.x && (!ok || gpfErr.Type != outsideLineRangeError) {
				left = linePnt
			}
		}

		if !math.IsInf(left.x, 0) {
			return getPointDistance(newPoint(x, y), left)
		}

		for _, val := range poly.lines {
			linePnt, _ := val.getPointFromY(y)
			if linePnt.x < left.x {
				left = linePnt
			}
		}
		return getPointDistance(newPoint(x, y), left)
	case Right:
		right := newPoint(math.Inf(-1), y)
		for _, val := range poly.lines {
			linePnt, err := val.getPointFromY(y)
			if gpfErr, ok := err.(lineGetPointFromError); linePnt.x > right.x && (!ok || gpfErr.Type != outsideLineRangeError) {
				right = linePnt
			}
		}

		if !math.IsInf(right.x, 0) {
			return getPointDistance(newPoint(x, y), right)
		}

		for _, val := range poly.lines {
			linePnt, _ := val.getPointFromY(y)
			if linePnt.x > right.x {
				right = linePnt
			}
		}
		return getPointDistance(newPoint(x, y), right)
	default:
		return 0, 0
	}
}

func (poly *polygon) onPoint(p *point) bool {

	if !p.onBounding(poly.bounds) {
		return false
	}

	ray := newLine(newPoint(math.Floor(p.x), math.Floor(p.y)),
		newPoint(math.Floor(p.x+(poly.bounds.end.x-poly.bounds.start.x)), math.Floor(p.y)))

	c := make(chan int, len(poly.lines))
	intersects := 0
	for _, val := range poly.lines {
		val := val
		go func() {
			intersectPnt := getLineIntersection(ray, val)
			if intersectPnt == nil {
				c <- 0
				return
			}

			intersectCnt := 0
			isOnVertex := false
			if intersectPnt.onPoint(val.start) && val.end.y < intersectPnt.y {
				intersectCnt++
				isOnVertex = true
			} else if intersectPnt.onPoint(val.end) && val.start.y < intersectPnt.y {
				intersectCnt++
				isOnVertex = true
			} else if intersectPnt.onPoint(val.end) || intersectPnt.onPoint(val.start) {
				isOnVertex = true
			}

			if !isOnVertex {
				intersectCnt++
			}

			c <- intersectCnt
		}()
	}

	for i := 0; i < len(poly.lines); i++ {
		intersects += <-c
	}

	if intersects%2 == 0 {
		return false
	}

	return true
}

func (poly *polygon) onBounding(b *bounding) bool {

	if !b.onBounding(poly.bounds) {
		return false
	}

	if poly.onPoint(b.start) || poly.onPoint(b.end) ||
		poly.onPoint(newPoint(b.start.x, b.end.y)) ||
		poly.onPoint(newPoint(b.end.x, b.start.y)) {
		return true
	}

	boundLines := b.getLines()
	for _, val := range boundLines {
		if poly.onLine(val) {
			return true
		}
	}

	return false
}

func (poly *polygon) onLine(l *line) bool {

	if !l.bounds.onBounding(poly.bounds) {
		return false
	}

	for _, val := range poly.lines {
		if l.onLine(val) {
			return true
		}
	}

	return false
}

func (poly *polygon) onPolygon(poly2 *polygon) bool {

	if !poly.bounds.onBounding(poly2.bounds) {
		return false
	}

	for _, val := range poly.lines {
		if poly2.onLine(val) {
			return true
		}
	}

	return false
}

func (l *line) onPolygon(poly *polygon) bool {

	return poly.onLine(l)
}

func (b *bounding) onPolygon(poly *polygon) bool {

	return poly.onBounding(b)
}

func (p *point) onPolygon(poly *polygon) bool {

	return poly.onPoint(p)
}
