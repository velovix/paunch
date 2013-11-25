package paunch

import (
	"testing"
)

func TestPointOnPoint(t *testing.T) {

	point := NewPoint(2.0, 3.0)
	point1 := NewPoint(2.0, 3.0)
	point2 := NewPoint(5.0, 4.0)

	if !isPointOnPoint(point, point1) {
		t.Errorf("isPointOnPoint returned false when true is expected")
	}

	if isPointOnPoint(point, point2) {
		t.Errorf("isPointOnPoint returned true when false is expected")
	}
}

func TestPointOnBounding(t *testing.T) {

	bounding := NewBounding(NewPoint(1.0, 1.0), NewPoint(5.0, 3.0))
	point1 := NewPoint(2.0, 2.0)
	point2 := NewPoint(10.0, 10.0)

	if !isPointOnBounding(point1, bounding) {
		t.Errorf("isPointOnBounding returned false when true is expected")
	}

	if isPointOnBounding(point2, bounding) {
		t.Errorf("isPointOnBounding returned true when false is expected")
	}
}

func TestBoundingOnBounding(t *testing.T) {

	bounding := NewBounding(NewPoint(1.0, 1.0), NewPoint(5.0, 3.0))
	bounding1 := NewBounding(NewPoint(2.0, 1.0), NewPoint(3.0, 2.0))
	bounding2 := NewBounding(NewPoint(10.0, 10.0), NewPoint(12.0, 13.0))

	if !isBoundingOnBounding(bounding, bounding1) {
		t.Errorf("isBoundingOnBounding returned false when true is expected")
	}

	if isBoundingOnBounding(bounding, bounding2) {
		t.Errorf("isBoundingOnBounding returned true when false is expected")
	}
}

func TestPointOnLine(t *testing.T) {

	point := NewPoint(2.0, 2.0)
	line1 := NewLine(NewPoint(0.0, 0.0), NewPoint(5.0, 5.0))
	line2 := NewLine(NewPoint(0.0, 0.0), NewPoint(5.0, 1.0))

	if !isPointOnLine(point, line1) {
		t.Errorf("isPointOnLine returned false when true is expected")
	}

	if isPointOnLine(point, line2) {
		t.Errorf("isPointOnLine returned true when false is expected")
	}
}

func TestBoundingOnLine(t *testing.T) {

	bounding := NewBounding(NewPoint(0.0, 0.0), NewPoint(6.0, 5.0))
	line1 := NewLine(NewPoint(1.0, 2.0), NewPoint(8.0, 3.0))
	line2 := NewLine(NewPoint(7.0, 7.0), NewPoint(10.0, 10.0))

	if !isBoundingOnLine(bounding, line1) {
		t.Errorf("isBoundingOnLine returned false when true is expected")
	}

	if isBoundingOnLine(bounding, line2) {
		t.Errorf("isBoundingOnLine returned true when false is expected")
	}
}

func TestLineOnLine(t *testing.T) {

	line := NewLine(NewPoint(0.0, 0.0), NewPoint(5.0, 5.0))
	line1 := NewLine(NewPoint(1.0, 2.0), NewPoint(6.0, 2.1))
	line2 := NewLine(NewPoint(0.0, 1.0), NewPoint(5.0, 6.0))

	if !isLineOnLine(line, line1) {
		t.Errorf("isLineOnLine returned false when true is expected")
	}

	if isLineOnLine(line, line2) {
		t.Errorf("isLineOnLine returned true when false is expected")
	}
}

func TestIsPointOnPolygon(t *testing.T) {

	polygon := NewPolygon([]Point{NewPoint(0.0, 0.0), NewPoint(3.0, 0.1), NewPoint(2.0, 4.0), NewPoint(0.0, 4.0)})
	point1 := NewPoint(1.0, 1.0)
	point2 := NewPoint(3.0, 4.0)

	if !isPointOnPolygon(point1, polygon) {
		t.Errorf("isPointOnPolygon returned false when true is expected")
	}

	if isPointOnPolygon(point2, polygon) {
		t.Errorf("isPointOnPolygon returned true when false is expected")
	}
}

func TestIsBoundingOnPolygon(t *testing.T) {

	polygon := NewPolygon([]Point{NewPoint(0.0, 0.0), NewPoint(3.0, 0.1), NewPoint(2.0, 4.0), NewPoint(0.0, 4.0)})
	bounding1 := NewBounding(NewPoint(1.0, 1.0), NewPoint(3.0, 2.0))
	bounding2 := NewBounding(NewPoint(3.0, 4.0), NewPoint(6.0, 6.0))

	if !isBoundingOnPolygon(bounding1, polygon) {
		t.Errorf("isBoundingOnPolygon returned false when true is expected")
	}

	if isBoundingOnPolygon(bounding2, polygon) {
		t.Errorf("isBoundingOnPolygon returned true when false is expected")
	}
}

func TestIsLineOnPolygon(t *testing.T) {

	polygon := NewPolygon([]Point{NewPoint(0.0, 0.0), NewPoint(3.0, 0.1), NewPoint(2.0, 4.0), NewPoint(0.0, 4.0)})
	line1 := NewLine(NewPoint(1.0, 1.0), NewPoint(3.0, 2.0))
	line2 := NewLine(NewPoint(3.0, 4.0), NewPoint(6.0, 6.0))

	if !isLineOnPolygon(line1, polygon) {
		t.Errorf("isLineOnPolygon returned false when true is expected")
	}

	if isLineOnPolygon(line2, polygon) {
		t.Errorf("isLineOnPolygon returned true when false is expected")
	}
}

func TestIsPolygonOnPolygon(t *testing.T) {

	polygon := NewPolygon([]Point{NewPoint(0.0, 0.0), NewPoint(3.0, 0.1), NewPoint(2.0, 4.0), NewPoint(0.0, 4.0)})
	polygon1 := NewPolygon([]Point{NewPoint(1.0, 1.0), NewPoint(3.0, 1.0), NewPoint(2.0, 7.0)})
	polygon2 := NewPolygon([]Point{NewPoint(7.0, 7.0), NewPoint(10.0, 7.0), NewPoint(8.0, 10.0)})

	if !isPolygonOnPolygon(polygon, polygon1) {
		t.Errorf("isPolygonOnPolygon returned false when true is expected")
	}

	if isPolygonOnPolygon(polygon, polygon2) {
		t.Errorf("isPolygonOnPolygon returned true when false is expected")
	}
}
