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

	var effect Effect
	effectList := []int{0, 0}
	effectList[0], err = effect.NewEffect(VERTEX, "basic")
	if err != nil {
		t.Errorf(".NewEffect(VERTEX, \"basic\") returned %s", err)
	}
	effectList[1], err = effect.NewEffect(FRAGMENT, "basic")
	if err != nil {
		t.Errorf(".NewEffect(FRAGMENT, \"basic\") returned %s", err)
	}

	var effectList_id int
	effectList_id, err = effect.NewEffectList(effectList)
	if err != nil {
		t.Errorf(".NewEffectList(effectList) returned %s", err)
	}

	effect.UseEffectList(effectList_id)
}
