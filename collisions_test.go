package paunch

import (
	"testing"
)

func TestPointOnPoint(t *testing.T) {

	point := NewPoint(2.0, 3.0)
	point1 := NewPoint(2.0, 3.0)
	point2 := NewPoint(5.0, 4.0)

	if !point.OnPoint(point1) {
		t.Errorf("point.OnPoint returned false when true is expected")
	}

	if point.OnPoint(point2) {
		t.Errorf("point.OnPoint returned true when false is expected")
	}
}

func TestPointOnBounding(t *testing.T) {

	bounding := NewBounding(NewPoint(1.0, 1.0), NewPoint(5.0, 3.0))
	point1 := NewPoint(2.0, 2.0)
	point2 := NewPoint(10.0, 10.0)

	if !bounding.OnPoint(point1) {
		t.Errorf("bounding.OnPoint returned false when true is expected")
	}

	if bounding.OnPoint(point2) {
		t.Errorf("bounding.OnPoint returned true when false is expected")
	}
}

func TestBoundingOnBounding(t *testing.T) {

	bounding := NewBounding(NewPoint(1.0, 1.0), NewPoint(5.0, 3.0))
	bounding1 := NewBounding(NewPoint(2.0, 1.0), NewPoint(3.0, 2.0))
	bounding2 := NewBounding(NewPoint(10.0, 10.0), NewPoint(12.0, 13.0))

	if !bounding.OnBounding(bounding1) {
		t.Errorf("bounding.OnBounding returned false when true is expected")
	}

	if bounding.OnBounding(bounding2) {
		t.Errorf("bounding.OnBounding returned true when false is expected")
	}
}

func TestPointOnLine(t *testing.T) {

	point := NewPoint(2.0, 2.0)
	line1 := NewLine(NewPoint(0.0, 0.0), NewPoint(5.0, 5.0))
	line2 := NewLine(NewPoint(0.0, 0.0), NewPoint(5.0, 1.0))

	if !point.OnLine(line1) {
		t.Errorf("point.OnLine returned false when true is expected")
	}

	if point.OnLine(line2) {
		t.Errorf("point.OnLine returned true when false is expected")
	}
}

func TestBoundingOnLine(t *testing.T) {

	bounding := NewBounding(NewPoint(0.0, 0.0), NewPoint(6.0, 5.0))
	line1 := NewLine(NewPoint(1.0, 2.0), NewPoint(8.0, 3.0))
	line2 := NewLine(NewPoint(7.0, 7.0), NewPoint(10.0, 10.0))

	if !bounding.OnLine(line1) {
		t.Errorf("bounding.OnLine returned false when true is expected")
	}

	if bounding.OnLine(line2) {
		t.Errorf("bounding.OnLine returned true when false is expected")
	}
}

func TestLineOnLine(t *testing.T) {

	line := NewLine(NewPoint(0.0, 0.0), NewPoint(5.0, 5.0))
	line1 := NewLine(NewPoint(1.0, 2.0), NewPoint(6.0, 2.1))
	line2 := NewLine(NewPoint(0.0, 1.0), NewPoint(5.0, 6.0))

	if !line.OnLine(line1) {
		t.Errorf("line.OnLine returned false when true is expected")
	}

	if line.OnLine(line2) {
		t.Errorf("line.OnLine returned true when false is expected")
	}
}

func TestIsPointOnPolygon(t *testing.T) {

	polygon := NewPolygon([]*Point{NewPoint(0.0, 0.0), NewPoint(3.0, 0.1), NewPoint(2.0, 4.0), NewPoint(0.0, 4.0)})
	point1 := NewPoint(1.0, 1.0)
	point2 := NewPoint(3.0, 4.0)

	if !polygon.OnPoint(point1) {
		t.Errorf("polygon.OnPoint returned false when true is expected")
	}

	if polygon.OnPoint(point2) {
		t.Errorf("polygon.OnPoint returned true when false is expected")
	}
}

func TestIsBoundingOnPolygon(t *testing.T) {

	polygon := NewPolygon([]*Point{NewPoint(0.0, 0.0), NewPoint(3.0, 0.1), NewPoint(2.0, 4.0), NewPoint(0.0, 4.0)})
	bounding1 := NewBounding(NewPoint(1.0, 1.0), NewPoint(3.0, 2.0))
	bounding2 := NewBounding(NewPoint(3.0, 4.0), NewPoint(6.0, 6.0))

	if !polygon.OnBounding(bounding1) {
		t.Errorf("polygon.OnBounding returned false when true is expected")
	}

	if polygon.OnBounding(bounding2) {
		t.Errorf("polygon.OnBounding returned true when false is expected")
	}
}

func TestIsLineOnPolygon(t *testing.T) {

	polygon := NewPolygon([]*Point{NewPoint(0.0, 0.0), NewPoint(3.0, 0.1), NewPoint(2.0, 4.0), NewPoint(0.0, 4.0)})
	line1 := NewLine(NewPoint(1.0, 1.0), NewPoint(3.0, 2.0))
	line2 := NewLine(NewPoint(3.0, 4.0), NewPoint(6.0, 6.0))

	if !polygon.OnLine(line1) {
		t.Errorf("polygon.OnLine returned false when true is expected")
	}

	if polygon.OnLine(line2) {
		t.Errorf("polygon.OnLine returned true when false is expected")
	}
}

func TestIsPolygonOnPolygon(t *testing.T) {

	polygon := NewPolygon([]*Point{NewPoint(0.0, 0.0), NewPoint(3.0, 0.1), NewPoint(2.0, 4.0), NewPoint(0.0, 4.0)})
	polygon1 := NewPolygon([]*Point{NewPoint(1.0, 1.0), NewPoint(3.0, 1.0), NewPoint(2.0, 7.0)})
	polygon2 := NewPolygon([]*Point{NewPoint(7.0, 7.0), NewPoint(10.0, 7.0), NewPoint(8.0, 10.0)})

	if !polygon.OnPolygon(polygon1) {
		t.Errorf("polygon.OnPolygon returned false when true is expected")
	}

	if polygon.OnPolygon(polygon2) {
		t.Errorf("polygon.OnPolygon returned true when false is expected")
	}
}

func TestShapes(t *testing.T) {

	collision := NewPolygon([]*Point{NewPoint(0.0, 0.0), NewPoint(3.0, 0.1), NewPoint(2.0, 4.0), NewPoint(0.0, 4.0)})
	collision1 := NewLine(NewPoint(1.0, 1.0), NewPoint(3.0, 2.0))
	collision2 := NewLine(NewPoint(3.0, 4.0), NewPoint(6.0, 6.0))

	if !Collides(collision, collision1) {
		t.Errorf("Collides returned false when true is expected")
	}

	if Collides(collision, collision2) {
		t.Errorf("Collides returned true when false is expected")
	}
}

func TestDistanceToTangentPointOnPoint(t *testing.T) {

	tangent := NewPoint(3.0, 5.0).DistanceToTangentPoint(NewPoint(5.0, 6.0), Left)

	if tangent.X != 2.0 || tangent.Y != 1.0 {
		t.Errorf("method DistanceToTangentPoint returns invalid values when used on a Point (%f, %f)", tangent.X, tangent.Y)
	}
}

func TestDistanceToTangentPointOnBounding(t *testing.T) {

	tangent := NewBounding(NewPoint(0.0, 0.0), NewPoint(3.0, 5.0)).DistanceToTangentPoint(NewPoint(2.0, 6.0), Up)

	if tangent.X != 0.0 || tangent.Y != -1.0 {
		t.Errorf("method DistanceToTangentPoint returns invalid values when used on a Bounding (%f, %f)", tangent.X, tangent.Y)
	}
}

func TestDistanceToTangentPointOnLine(t *testing.T) {

	tangent := NewLine(NewPoint(0.0, 0.0), NewPoint(3.0, 3.0)).DistanceToTangentPoint(NewPoint(1.0, 2.0), Up)

	if tangent.X != 0.0 || tangent.Y != -1.0 {
		t.Errorf("method DistanceToTangentPoint returns invalid values when used on a Line (%f, %f)", tangent.X, tangent.Y)
	}
}

func TestDistanceToTangentPointOnPolygon(t *testing.T) {

	polygon := NewPolygon([]*Point{NewPoint(0.0, 0.0), NewPoint(5.0, 1.0), NewPoint(4.0, 3.0), NewPoint(-1.0, 2.0)})
	tangent := polygon.DistanceToTangentPoint(NewPoint(3.9, 6.0), Up)

	if tangent.X != 0.0 || tangent.Y != -5.22 {
		t.Errorf("method DistanceToTangentPoint returns invalid values when used on a Polygon (%f, %f)", tangent.X, tangent.Y)
	}
}
