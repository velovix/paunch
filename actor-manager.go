package paunch

// ActorManager triggers methods with the On- prefix when appropriate given the
// Actors supplied to it.
type ActorManager struct {
	actors []Actor
}

// NewActorManager creates a new ActorManager.
func NewActorManager() ActorManager {

	return ActorManager{make([]Actor, 0)}
}

// SetActors sets the Actors used by the ActorManager object to the given
// slice. The ActorManager does not make a seperate copy of the slice, so
// changes to the slice will affect the ActorManager.
func (actorManager *ActorManager) SetActors(actors []Actor) {

	actorManager.actors = actors
}

// GetActors returns the slice of Actor objects the ActorManager has stored.
func (actorManager *ActorManager) GetActors() []Actor {

	return actorManager.actors
}

// RunKeyEvent simulates a key event, triggering the expected response from
// the ActorManager object's Actors.
func (actorManager ActorManager) RunKeyEvent(key Key, action Action) {

	for i := range actorManager.actors {
		if val, ok := actorManager.actors[i].(ActorKeyboarder); ok {
			val.OnKeyboard(key, action)
		}
	}
}

// RunMouseButtonEvent simulates a mouse button event, triggering the
// expected response from the ActorManager object's Actors.
func (actorManager ActorManager) RunMouseButtonEvent(button MouseButton, action Action, x, y int) {

	for i := range actorManager.actors {
		if val, ok := actorManager.actors[i].(ActorMouseButtoner); ok {
			val.OnMouseButton(button, action, x, y)
		}
	}
}

// RunMousePositionEvent simulates a mouse position event, triggering the
// expected response from the ActorManager object's Actors.
func (actorManager ActorManager) RunMousePositionEvent(x, y int) {

	for i := range actorManager.actors {
		if val, ok := actorManager.actors[i].(ActorMousePositioner); ok {
			val.OnMousePosition(x, y)
		}
	}
}

// RunMouseEnterWindowEvent simulates a mouse enter window event, triggering
// the expected response from the ActorManager object's Actors.
func (actorManager ActorManager) RunMouseEnterWindowEvent(x, y int, entered bool) {

	for i := range actorManager.actors {
		if val, ok := actorManager.actors[i].(ActorMouseEnterWindower); ok {
			val.OnMouseEnterWindow(x, y, entered)
		}
	}
}

// RunWindowFocusEvent simulates a window focus event, triggering the
// expected response from the ActorManager object's Actors.
func (actorManager ActorManager) RunWindowFocusEvent(focused bool) {

	for i := range actorManager.actors {
		if val, ok := actorManager.actors[i].(ActorWindowFocuser); ok {
			val.OnWindowFocus(focused)
		}
	}
}

// RunWindowResizeEvent simulates a window resize event, triggering the
// expected response from the ActorManager object's Actors.
func (actorManager ActorManager) RunWindowResizeEvent(width, height int) {

	for i := range actorManager.actors {
		if val, ok := actorManager.actors[i].(ActorWindowResizer); ok {
			val.OnWindowResize(width, height)
		}
	}
}

// RunJoystickButtonEvent simulates a joystick button event, triggering the
// expected response from the ActorManager object's Actors.
func (actorManager ActorManager) RunJoystickButtonEvent(button int, action Action) {

	for i := range actorManager.actors {
		if val, ok := actorManager.actors[i].(ActorJoystickButtoner); ok {
			val.OnJoystickButton(button, action)
		}
	}
}

// RunJoystickAxisEvent simulates a joystick axis event, triggering the
// expected response from the ActorManager object's Actors.
func (actorManager ActorManager) RunJoystickAxisEvent(device int, value float64) {

	for i := range actorManager.actors {
		if val, ok := actorManager.actors[i].(ActorJoystickAxiser); ok {
			val.OnJoystickAxis(device, value)
		}
	}
}

// RunCollisionEvent checks for collisions between the ActorManager's Actors
// and triggers appropriate methods.
func (actorManager ActorManager) RunCollisionEvent() {

	for i := range actorManager.actors {
		actorCollider, ok := actorManager.actors[i].(ActorCollider)
		if !ok {
			continue
		}
		colliders1 := actorCollider.GetColliders()

		for _, val := range actorManager.actors {
			if actorManager.actors[i] == val {
				continue
			}

			actorCollider2, ok2 := val.(ActorCollider)
			if !ok2 {
				continue
			}
			colliders2 := actorCollider2.GetColliders()

			for _, col1 := range colliders1 {
				for _, col2 := range colliders2 {
					if Collides(col1, col2) {
						actorCollider.OnCollision(col1, col2, val)
					}
				}
			}
		}
	}
}

// RunTickEvent runs a tick event, triggering the expected response from
// the ActorManager object's Actors.
func (actorManager ActorManager) RunTickEvent() {

	for i := range actorManager.actors {
		if val, ok := actorManager.actors[i].(ActorTicker); ok {
			val.OnTick()
		}
	}
}

// RunDrawEvent runs a draw event, triggering the expected response from
// the ActorManager object's Actors.
func (actorManager ActorManager) RunDrawEvent() {

	for i := range actorManager.actors {
		if val, ok := actorManager.actors[i].(ActorDrawer); ok {
			val.OnDraw()
		}
	}
}

// Collides checks if the supplied Collider collides with any of the
// ActorManager's Actors.
func (actorManager ActorManager) Collides(collider Collider) bool {

	for i := range actorManager.actors {
		actorCollider, ok := actorManager.actors[i].(ActorCollider)
		if !ok {
			continue
		}

		collisions := actorCollider.GetColliders()
		for _, val := range collisions {
			if Collides(collider, val) {
				return true
			}
		}
	}

	return false
}
