package paunch

import (
	"testing"
)

var draw Draw

func TestInitDraw(t *testing.T) {

	var window Window
	err := window.Open(640, 480, "Test")
	if err != nil {
		t.Errorf(".Open(640, 480, \"Test\") returned %s", err)
	}

	err = draw.Init(window)
	if err != nil {
		t.Errorf("InitDraw() returned %s", err)
	}
}

func TestDrawTriangles(t *testing.T) {

	triangle := []float32{
		0.0, 0.0, 0.0,
		1.0, 0.0, 0.0,
		0.0, 1.0, 0.0}

	renderable, err := draw.NewRenderable(TRIANGLES, triangle)
	if err != nil {
		t.Errorf(".NewRenderable(TRIANGLE, triangle) returned %s", err)
	}

	draw.DrawRenderable(renderable)
}
