package paunch

// EventManager triggers methods with the On- prefix when appropriate given the
// objects supplied to it.
type EventManager struct {
	objects []interface{}
}

// NewEventManager creates a new EventManager.
func NewEventManager() EventManager {

	return EventManager{make([]interface{}, 0)}
}

// SetActors sets the objects used by the EventManager object to the given
// slice. The EventManager does not make a seperate copy of the slice, so
// changes to the slice will affect the EventManager.
func (eventManager *EventManager) SetObjects(objects []interface{}) {

	eventManager.objects = objects
}

// GetActors returns the slice of objects the EventManager has stored.
func (eventManager *EventManager) GetObjects() []interface{} {

	return eventManager.objects
}

// RunKeyEvent simulates a key event, triggering the expected response from
// the EventManager's objects.
func (eventManager EventManager) RunKeyEvent(key Key, action Action) {

	for i := range eventManager.objects {
		if val, ok := eventManager.objects[i].(ActorKeyboarder); ok {
			val.OnKeyboard(key, action)
		}
	}
}

// RunMouseButtonEvent simulates a mouse button event, triggering the
// expected response from the EventManager's objects.
func (eventManager EventManager) RunMouseButtonEvent(button MouseButton, action Action, x, y int) {

	for i := range eventManager.objects {
		if val, ok := eventManager.objects[i].(ActorMouseButtoner); ok {
			val.OnMouseButton(button, action, x, y)
		}
	}
}

// RunMousePositionEvent simulates a mouse position event, triggering the
// expected response from the EventManager's objects.
func (eventManager EventManager) RunMousePositionEvent(x, y int) {

	for i := range eventManager.objects {
		if val, ok := eventManager.objects[i].(ActorMousePositioner); ok {
			val.OnMousePosition(x, y)
		}
	}
}

// RunMouseEnterWindowEvent simulates a mouse enter window event, triggering
// the expected response from the EventManager's objects.
func (eventManager EventManager) RunMouseEnterWindowEvent(x, y int, entered bool) {

	for i := range eventManager.objects {
		if val, ok := eventManager.objects[i].(ActorMouseEnterWindower); ok {
			val.OnMouseEnterWindow(x, y, entered)
		}
	}
}

// RunWindowFocusEvent simulates a window focus event, triggering the
// expected response from the EventManager's objects.
func (eventManager EventManager) RunWindowFocusEvent(focused bool) {

	for i := range eventManager.objects {
		if val, ok := eventManager.objects[i].(ActorWindowFocuser); ok {
			val.OnWindowFocus(focused)
		}
	}
}

// RunWindowResizeEvent simulates a window resize event, triggering the
// expected response from the EventManager's objects.
func (eventManager EventManager) RunWindowResizeEvent(width, height int) {

	for i := range eventManager.objects {
		if val, ok := eventManager.objects[i].(ActorWindowResizer); ok {
			val.OnWindowResize(width, height)
		}
	}
}

// RunJoystickButtonEvent simulates a joystick button event, triggering the
// expected response from the EventManager's objects.
func (eventManager EventManager) RunJoystickButtonEvent(button int, action Action) {

	for i := range eventManager.objects {
		if val, ok := eventManager.objects[i].(ActorJoystickButtoner); ok {
			val.OnJoystickButton(button, action)
		}
	}
}

// RunJoystickAxisEvent simulates a joystick axis event, triggering the
// expected response from the EventManager's objects.
func (eventManager EventManager) RunJoystickAxisEvent(device int, value float64) {

	for i := range eventManager.objects {
		if val, ok := eventManager.objects[i].(ActorJoystickAxiser); ok {
			val.OnJoystickAxis(device, value)
		}
	}
}

// RunCollisionEvent checks for collisions between the EventManager's Actors
// and triggers appropriate methods.
func (eventManager EventManager) RunCollisionEvent() {

	for i := range eventManager.objects {
		actorCollider, ok := eventManager.objects[i].(ActorCollider)
		if !ok {
			continue
		}
		colliders1 := actorCollider.GetColliders()

		for _, val := range eventManager.objects {
			if eventManager.objects[i] == val {
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
// the EventManager's objects.
func (eventManager EventManager) RunTickEvent() {

	for i := range eventManager.objects {
		if val, ok := eventManager.objects[i].(ActorTicker); ok {
			val.OnTick()
		}
	}
}

// RunDrawEvent runs a draw event, triggering the expected response from
// the EventManager's objects.
func (eventManager EventManager) RunDrawEvent() {

	for i := range eventManager.objects {
		if val, ok := eventManager.objects[i].(ActorDrawer); ok {
			val.OnDraw()
		}
	}
}

// Collides checks if the supplied Collider collides with any of the
// EventManager's objects.
func (eventManager EventManager) Collides(collider Collider) bool {

	for i := range eventManager.objects {
		actorCollider, ok := eventManager.objects[i].(ActorCollider)
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
