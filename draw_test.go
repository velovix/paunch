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

	texCoords := []float32{
		0.0, 0.0,
		1.0, 0.0,
		0.0, 1.0}

	renderable, err := NewRenderable(TRIANGLES, triangle, texCoords)
	if err != nil {
		t.Errorf(".NewRenderable(TRIANGLE, triangle) returned %s", err)
	}

	DrawRenderable(renderable)

	var effect Effect
	effect.Init()
	err = effect.NewEffect(VERTEX, "basic")
	if err != nil {
		t.Errorf(".NewEffect(VERTEX, \"basic\") returned %s", err)
	}
	err = effect.NewEffect(FRAGMENT, "white")
	if err != nil {
		t.Errorf(".NewEffect(FRAGMENT, \"white\") returned %s", err)
	}

	err = effect.NewEffectList("simple", []string{"basic", "white"})
	if err != nil {
		t.Errorf(".NewEffectList(\"simple\", []string{\"basic\", \"white\"}) returned %s", err)
	}

	effect.UseEffectList("simple")
}
