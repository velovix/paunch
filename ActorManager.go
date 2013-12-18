package paunch

import (
	"fmt"
)

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
	fmt.Println(actorManager.actors)
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

			fmt.Println(actorManager.actors)
			return true
		}
	}

	fmt.Println(actorManager.actors)
	return false
}

// Run calls all relevant methods of the Actors supplied to the ActorManager.
func (actorManager ActorManager) Run() {

	for _, val := range actorManager.actors {

		val.Draw()
	}
}
