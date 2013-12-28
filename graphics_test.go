package paunch

import (
	"testing"
)

func TestOpenWindow(t *testing.T) {

	var window Window
	err := window.Open(640, 480, "Test")
	if err != nil {
		t.Errorf(".Open(640, 480, \"Test\") returned %s", err)
	}

	err = window.Destroy()
	if err != nil {
		t.Errorf(".Destroy() returned %s", err)
	}
}

func TestUpdateDisplay(t *testing.T) {

	var window Window
	err := window.Open(640, 480, "Test")
	if err != nil {
		t.Errorf(".Open(640, 480, \"Test\") returned %s", err)
	}

	err = window.UpdateDisplay()
	if err != nil {
		t.Errorf(".UpdateDisplay() returned %s", err)
	}
}

func TestUpdateEvents(t *testing.T) {

	var window Window
	err := window.Open(640, 480, "Test")
	if err != nil {
		t.Errorf(".Open(640, 480, \"Test\") returned %s", err)
	}

	err = window.UpdateEvents()
	if err != nil {
		t.Errorf(".UpdateEvents() returned %s", err)
	}
}

func TestInitDraw(t *testing.T) {

	var window Window
	err := window.Open(640, 480, "Test")
	if err != nil {
		t.Errorf("window.Open(640, 480, \"Test\") returned %s", err)
	}

	err = InitDraw(window)
	if err != nil {
		t.Errorf("draw.Init() returned %s", err)
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

func TestEffects(t *testing.T) {

	var effect Effect
	err := effect.Init()
	if err != nil {
		t.Errorf("effect.Init() returned %s", err)
	}

	err = effect.New("texture", "texture/")
	if err != nil {
		t.Errorf("effect.New(\"texture\", \"texture/\") returned %s", err)
	}

	err = effect.Use("texture")
	if err != nil {
		t.Errorf("effect.Use(\"texture\") returned %s", err)
	}
}
