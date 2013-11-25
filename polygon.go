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

func isPointOnPolygon(point Point, polygon Polygon) bool {

	if !isPointOnBounding(point, polygon.bounds) {
		return false
	}

	ray := NewLine(point, NewPoint(point.X+999, point.Y))

	intersects := 0
	for _, val := range polygon.lines {
		if isLineOnLine(ray, val) {
			intersects++
		}
	}

	if intersects%2 == 0 {
		return false
	} else {
		return true
	}
}

func isBoundingOnPolygon(bounding Bounding, polygon Polygon) bool {

	if !isBoundingOnBounding(bounding, polygon.bounds) {
		return false
	}

	if isPointOnPolygon(bounding.Start, polygon) || isPointOnPolygon(bounding.End, polygon) ||
		isPointOnPolygon(NewPoint(bounding.Start.X, bounding.End.Y), polygon) ||
		isPointOnPolygon(NewPoint(bounding.End.X, bounding.Start.Y), polygon) {
		return true
	}

	boundLines := bounding.getLines()
	for _, val := range boundLines {
		if isLineOnPolygon(val, polygon) {
			return true
		}
	}

	return false
}

func isLineOnPolygon(line Line, polygon Polygon) bool {

	if !isBoundingOnBounding(line.bounds, polygon.bounds) {
		return false
	}

	for _, val := range polygon.lines {
		if isLineOnLine(line, val) {
			return true
		}
	}

	return false
}

func isPolygonOnPolygon(polygon1, polygon2 Polygon) bool {

	if !isBoundingOnBounding(polygon1.bounds, polygon2.bounds) {
		return false
	}

	for _, val := range polygon1.lines {
		if isLineOnPolygon(val, polygon2) {
			return true
		}
	}

	return false
}
