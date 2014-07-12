// window.go
// by Tyler Compton
// This example shows how to get a blank window on screen using Paunch.

package main

import (
	"github.com/velovix/paunch"
)

func main() {

	// Sets some parameters for the window
	paunch.SetWindowSize(640, 480)
	paunch.SetWindowTitle("Test Window")

	// Starts Paunch, initializing things and opening the window
	err := paunch.Start(paunch.VersionAutomatic)
	if err != nil {
		panic(err)
	}
	defer paunch.Stop()

	// Runs for as long as Paunch does not recieve a "close event" (i.e. if the
	// user attempts to close the window)
	for !paunch.ShouldClose() {

		// Updates the display and checks for events
		paunch.UpdateDisplay()
		paunch.UpdateEvents()
	}
}
