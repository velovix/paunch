package paunch

import (
	"testing"
)

func TestOpenWindow(t *testing.T) {

	err := InitWindows()
	if err != nil {
		t.Errorf("InitWindows() returned %s", err)
	}

	window := NewWindow(640, 480, 640, 480, false, "Test")
	err = window.Open()
	if err != nil {
		t.Errorf(".Open(640, 480, 640, 480, false, \"Test\") returned %s", err)
	}

	err = window.Close()
	if err != nil {
		t.Errorf(".Close() returned %s", err)
	}
}

func TestUpdateDisplay(t *testing.T) {

	err := InitWindows()
	if err != nil {
		t.Errorf("InitWindows() returned %s", err)
	}

	window := NewWindow(640, 480, 640, 480, false, "Test")
	err = window.Open()
	if err != nil {
		t.Errorf(".Open(640, 480, 640, 480, false, \"Test\") returned %s", err)
	}

	err = window.UpdateDisplay()
	if err != nil {
		t.Errorf(".UpdateDisplay() returned %s", err)
	}
}

func TestUpdateEvents(t *testing.T) {

	err := InitWindows()
	if err != nil {
		t.Errorf("InitWindows() returned %s", err)
	}

	window := NewWindow(640, 480, 640, 480, false, "Test")
	err = window.Open()
	if err != nil {
		t.Errorf(".Open(640, 480, 640, 480, false, \"Test\") returned %s", err)
	}

	err = window.UpdateEvents()
	if err != nil {
		t.Errorf(".UpdateEvents() returned %s", err)
	}
}

func TestInitDraw(t *testing.T) {

	err := InitWindows()
	if err != nil {
		t.Errorf("InitWindows() returned %s", err)
	}

	window := NewWindow(640, 480, 640, 480, false, "Test")
	err = window.Open()
	if err != nil {
		t.Errorf("window.Open(640, 480, 640, 480, false, \"Test\") returned %s", err)
	}

	err = InitDraw()
	if err != nil {
		t.Errorf("InitDraw() returned %s", err)
	}
}

func TestDrawRenderable(t *testing.T) {

	triangle := []float64{
		0.0, 0.0,
		1.0, 0.0,
		0.0, 1.0}

	renderable, err := NewRenderable(Triangles, triangle)
	if err != nil {
		t.Errorf("NewRenderable(TRIANGLE, triangle) returned %s", err)
	}

	err = renderable.Draw(0)
	if err != nil {
		t.Errorf("renderable.Draw(0) returned %s", err)
	}
}

func TestDrawTexturedRenderable(t *testing.T) {

	triangles := []float64{
		0.0, 0.0,
		1.0, 0.0,
		0.0, 1.0,
		1.0, 1.0,
		1.0, 0.0,
		0.0, 1.0}

	texCoords := []float64{
		0.0, 0.0,
		1.0, 0.0,
		0.0, 1.0,
		1.0, 1.0,
		1.0, 0.0,
		0.0, 1.0}

	renderable, err := NewRenderable(Triangles, triangles)
	if err != nil {
		t.Errorf("NewRenderable(TRIANGLE, triangles) returned %s", err)
	}

	err = renderable.Texture(texCoords, "img/test.png", 2)
	if err != nil {
		t.Errorf("renderable.Texture(texCoords, \"img/test.png\", 2) returned %s", err)
	}

	err = renderable.Draw(0)
	if err != nil {
		t.Errorf("renderable.Draw(0) returned %s", err)
	}
}

func TestDrawSurfaceRenderable(t *testing.T) {

	renderable, err := NewRenderableSurface(0, 0, 1, 1, "img/test.png", 2)
	if err != nil {
		t.Errorf("NewRenderable(0, 0, 1, 1 \"img/test.png\", 2) returned %s", err)
	}

	err = renderable.Draw(0)
	if err != nil {
		t.Errorf("renderable.Draw(0) returned %s", err)
	}
}

var vertex = `#version 330
layout(location = 0) in vec4 position;

void main() {
	vec4 vertex = position;
	gl_Position = vertex;
}`

var fragment = `#version 330
out vec4 fragColor;

void main() {
	fragColor = vec4(1.0, 0.0, 0.0, 1.0);
}`

func TestEffects(t *testing.T) {

	effect, err := NewEffectFromStrings([]string{vertex, fragment}, []ShaderType{Vertex, Fragment})
	if err != nil {
		t.Errorf("NewEffect([]string{vertex, fragment}, []ShaderType{Vertex, Shader}) returned %s", err)
	}

	err = UseEffect(&effect)
	if err != nil {
		t.Errorf("UseEffect(&effect) returned %s", err)
	}
}
