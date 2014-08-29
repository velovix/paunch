// event-manager.go
// by Tyler Compton
// This example shows how to capture user input and draw objects using the
// EventManager object.

package main

import (
	"github.com/velovix/paunch"
	"time"
)

var (
	eventManager *paunch.EventManager
	effect       *paunch.Effect
)

// A game object. The Player object will display test.png and be movable with
// the arrow keys.
type Player struct {
	image *paunch.Sprite
}

func NewPlayer(x, y float64) Player {

	var player Player
	var err error

	player.image, err = paunch.NewSpriteFromImage(x, y, "./test.png", 1)
	if err != nil {
		panic(err)
	}

	return player
}

// This method will be called by the EventManager object when a keyboard key
// is pressed.
func (player *Player) OnKeyboard(key paunch.Key, action paunch.Action) {

	if key == paunch.KeyLeft && action == paunch.Hold {
		player.image.Move(-3, 0)
	} else if key == paunch.KeyRight && action == paunch.Hold {
		player.image.Move(3, 0)
	} else if key == paunch.KeyDown && action == paunch.Hold {
		player.image.Move(0, -3)
	} else if key == paunch.KeyUp && action == paunch.Hold {
		player.image.Move(0, 3)
	}
}

// This method will be called by the EventManager when a draw event happens.
func (player *Player) OnDraw() {

	player.image.Draw(0)
}

func main() {

	paunch.SetWindowSize(640, 480)
	paunch.SetWindowTitle("Use the arrow keys to move the object!")

	err := paunch.Start(paunch.VersionAutomatic)
	if err != nil {
		panic(err)
	}
	defer paunch.Stop()

	effect, err = paunch.NewEffectFromDirectory("./shader/")
	if err != nil {
		panic(err)
	}
	paunch.UseEffect(effect)

	player := NewPlayer(288, 208)

	eventManager = paunch.NewEventManager()
	eventManager.Objects = []interface{}{&player} // Add the Player object to the EventManager's object list.
	eventManager.GetUserEvents(true)              // Set the EventManager to automatically respond to user events.

	lastFrame := time.Now()

	for !paunch.ShouldClose() {

		effect.SetVariable2f("screen_size", 640, 480)
		effect.SetVariablei("tex_id", 0)

		paunch.Clear()

		// Only check for events 1/60th of a second. This operation is not
		// Paunch-specific, but effectively limits the physics framerate to
		// 60 FPS.
		if time.Since(lastFrame).Seconds() >= (1.0 / 60.0) {
			lastFrame = time.Now()
			paunch.UpdateEvents()
		}

		// Has the EventManager run a draw event, which calls the OnDraw
		// methods of it's objects.
		eventManager.RunDrawEvent()

		paunch.UpdateDisplay()
	}
}
