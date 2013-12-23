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

func (actorManager ActorManager) keyEvent(key, action int) {

	for i := range actorManager.actors {
		if val, ok := actorManager.actors[i].(ActorKeyboarder); ok {
			val.OnKeyboard(key, action)
		}
	}
}

func (actorManager ActorManager) mouseButtonEvent(button, action, x, y int) {

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

func checkActorCollisions(actor1, actor2 ActorCollider) (bool, Collider, Collider) {

	c1 := actor1.GetColliders()
	c2 := actor2.GetColliders()

	for _, val := range c1 {
		for _, val2 := range c2 {
			if Collides(val, val2) {
				return true, val, val2
			}
		}
	}

	return false, nil, nil
}

// Tick runs all non-graphics tasks on the ActorManager's Actors.
func (actorManager ActorManager) Tick() {

	for i, val := range actorManager.actors {

		if colliderVal, ok := val.(ActorCollider); ok {
			for j, val2 := range actorManager.actors {
				if i == j {
					continue
				}

				if colliderVal2, ok2 := val2.(ActorCollider); ok2 {
					if collided, c1, c2 := checkActorCollisions(colliderVal, colliderVal2); collided {
						colliderVal.OnCollision(c1, c2, colliderVal2)
					}
				}
			}
		}

		if ticker, ok := val.(ActorTicker); ok {
			ticker.OnTick()
		}
	}
}

// Draw runs all graphics-related tasks on the ActorManager's Actors.
func (actorManager ActorManager) Draw() {

	for _, val := range actorManager.actors {
		if drawer, ok := val.(ActorDrawer); ok {
			drawer.OnDraw()
		}
	}
}
