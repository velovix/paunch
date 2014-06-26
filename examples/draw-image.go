// draw-image.go
// by Tyler Compton
// This example shows how to draw an image onto the screen.

package main

import (
	"github.com/velovix/paunch"
)

var (
	image  *paunch.Renderable
	effect *paunch.Effect
)

func main() {

	paunch.SetWindowSize(640, 480)
	paunch.SetWindowTitle("Test Window")

	err := paunch.Start()
	if err != nil {
		panic(err)
	}
	defer paunch.Stop()

	// Create an Effect object using the GLSL shader files found in the shader
	// directory
	effect, err = paunch.NewEffectFromDirectory("./shader/")
	if err != nil {
		panic(err)
	}
	paunch.UseEffect(effect) // Set the Effect object to be used

	// Create a new Renderable object for drawing test.png at x=288 y=208
	image, err = paunch.NewRenderableFromImage(288, 208, "./test.png", 1)
	if err != nil {
		panic(err)
	}

	for !paunch.ShouldClose() {

		// Send the necessary parameters to the GLSL shaders before drawing
		effect.SetVariable2f("screen_size", 640, 480)
		effect.SetVariablei("tex_id", 0)

		// Clear the screen of the last frame so that the new frame can be
		// drawn.
		paunch.Clear()

		// Draws the Renderable object
		image.Draw(0)

		paunch.UpdateDisplay()
		paunch.UpdateEvents()
	}
}
