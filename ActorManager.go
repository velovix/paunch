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

func checkActorCollisions(actor1, actor2 Actor) (bool, Collider, Collider) {

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

// Run calls all relevant methods of the Actors supplied to the ActorManager.
func (actorManager ActorManager) Run() {

	for i, val := range actorManager.actors {

		for j, val2 := range actorManager.actors {
			if i == j {
				continue
			}
			if ok, c1, c2 := checkActorCollisions(val, val2); ok {
				val.OnCollision(c1, c2, val2)
			}
		}

		val.Draw()
	}
}
