package paunch

import (
	"testing"
)

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

func TestDrawTriangles(t *testing.T) {

	triangle := []float32{
		0.0, 0.0,
		1.0, 0.0,
		0.0, 1.0}

	/*texCoords := []float32{
	0.0, 0.0,
	1.0, 0.0,
	0.0, 1.0}*/

	renderable, err := NewRenderable(TRIANGLES, triangle)
	if err != nil {
		t.Errorf(".NewRenderable(TRIANGLE, triangle) returned %s", err)
	}

	var effect Effect
	err = effect.Init()
	if err != nil {
		t.Errorf(".Init() returned %s", err)
	}

	err = effect.New("texture", "texture/")
	if err != nil {
		t.Errorf(".New(\"texture\", \"texture\") returned %s", err)
	}
	err = effect.Use("texture")
	if err != nil {
		t.Errorf(".New(\"texture\", \"texture\") returned %s", err)
	}

	/*err = renderable.Texture(texCoords, "img/test.png")
	if err != nil {
		t.Errorf(".Texture(texCoords, \"img/test.png\") returned %s", err)
	}*/

	renderable.Draw()
}
