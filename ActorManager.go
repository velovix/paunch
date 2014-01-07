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

// Add adds a new Actor to the ActorManager.
func (actorManager *ActorManager) Add(actor Actor) {

	actorManager.actors = append(actorManager.actors, actor)
}

// Remove removes all instances of the supplied Actor from the ActorManager.
func (actorManager *ActorManager) Remove(actor Actor) bool {

	for i := range actorManager.actors {
		if actor == actorManager.actors[i] {
			temp := make([]Actor, len(actorManager.actors)-1)

			for j, val := range actorManager.actors {
				if j != i {
					temp = append(temp, val)
				}
			}

			actorManager.actors = temp

			return true
		}
	}

	return false
}

func (actorManager ActorManager) keyEvent(key Key, action Action) {

	for i := range actorManager.actors {
		if val, ok := actorManager.actors[i].(ActorKeyboarder); ok {
			val.OnKeyboard(key, action)
		}
	}
}

func (actorManager ActorManager) mouseButtonEvent(button MouseButton, action Action, x, y int) {

	for i := range actorManager.actors {
		if val, ok := actorManager.actors[i].(ActorMouseButtoner); ok {
			val.OnMouseButton(button, action, x, y)
		}
	}
}

func (actorManager ActorManager) mousePositionEvent(x, y int) {

	for i := range actorManager.actors {
		if val, ok := actorManager.actors[i].(ActorMousePositioner); ok {
			val.OnMousePosition(x, y)
		}
	}
}

func (actorManager ActorManager) mouseEnterWindowEvent(x, y int, entered bool) {

	for i := range actorManager.actors {
		if val, ok := actorManager.actors[i].(ActorMouseEnterWindower); ok {
			val.OnMouseEnterWindow(x, y, entered)
		}
	}
}

func (actorManager ActorManager) windowFocusEvent(focused bool) {

	for i := range actorManager.actors {
		if val, ok := actorManager.actors[i].(ActorWindowFocuser); ok {
			val.OnWindowFocus(focused)
		}
	}
}

func (actorManager ActorManager) joystickButtonEvent(button int, action Action) {

	for i := range actorManager.actors {
		if val, ok := actorManager.actors[i].(ActorJoystickButtoner); ok {
			val.OnJoystickButton(button, action)
		}
	}
}

func (actorManager ActorManager) runCollisions(actor Actor) {

	actorCollider, ok := actor.(ActorCollider)
	if !ok {
		return
	}
	colliders1 := actorCollider.GetColliders()

	for _, val := range actorManager.actors {
		if actor == val {
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

func (actorManager ActorManager) runTicks(actor Actor) {

	actorTicker, ok := actor.(ActorTicker)
	if !ok {
		return
	}
	actorTicker.OnTick()
}

func (actorManager ActorManager) runDraws(actor Actor) {

	actorDrawer, ok := actor.(ActorDrawer)
	if !ok {
		return
	}
	actorDrawer.OnDraw()
}

// Tick runs all non-graphics tasks on the ActorManager's Actors.
func (actorManager ActorManager) Tick() {

	for _, val := range actorManager.actors {

		actorManager.runCollisions(val)
		actorManager.runTicks(val)
	}
}

// Draw runs all graphics-related tasks on the ActorManager's Actors.
func (actorManager ActorManager) Draw() {

	for _, val := range actorManager.actors {
		actorManager.runDraws(val)
	}
}
