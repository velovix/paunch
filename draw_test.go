package paunch

import (
	"testing"
)

func TestInitDraw(t *testing.T) {

	var window Window
	err := window.Open(640, 480, "Test")
	if err != nil {
		t.Errorf(".Open(640, 480, \"Test\") returned %s", err)
	}

	err = InitDraw(window)
	if err != nil {
		t.Errorf("InitDraw() returned %s", err)
	}
}

func TestDrawTriangles(t *testing.T) {

	triangle := []DrawFloat{
		0.0, 0.0, 0.0,
		1.0, 0.0, 0.0,
		0.0, 1.0, 0.0}

	Draw(TRIANGLES, triangle)
}
