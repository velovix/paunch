package paunch

type Polygon struct {
	lines  []Line
	bounds Bounding
}

func NewPolygon(points []Point) Polygon {

	var polygon Polygon
	polygon.lines = make([]Line, len(points))

	min := points[0]
	max := points[0]

	for i := 1; i < len(points); i++ {
		polygon.lines[i] = NewLine(points[i-1], points[i])

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
	polygon.lines[len(points)-1] = NewLine(points[len(points)-1], points[0])

	polygon.bounds = NewBounding(min, max)

	return polygon
}

func (polygon Polygon) OnPoint(point Point) bool {

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

func (polygon Polygon) OnBounding(bounding Bounding) bool {

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

func (polygon Polygon) OnLine(line Line) bool {

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

func (polygon1 Polygon) OnPolygon(polygon2 Polygon) bool {

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

func (line Line) OnPolygon(polygon Polygon) bool {

	return polygon.OnLine(line)
}

func (bounding Bounding) OnPolygon(polygon Polygon) bool {

	return polygon.OnBounding(bounding)
}

func (point Point) OnPolygon(polygon Polygon) bool {

	return polygon.OnPoint(point)
}
