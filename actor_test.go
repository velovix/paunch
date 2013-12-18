package paunch

import (
	"testing"
)

type testActorObject struct {
	id int
}

func (obj *testActorObject) Draw() {

}

func TestActorManager(t *testing.T) {

	var actorManager ActorManager

	var test1 testActorObject = testActorObject{1}
	var test2 testActorObject = testActorObject{2}

	actorManager.Add(&test1)
	actorManager.Add(&test2)

	if !actorManager.Remove(&test1) {
		t.Errorf("could not remove Actor from ActorManager")
	}

	if actorManager.Remove(&test1) {
		t.Errorf("double remove of Actor from ActorManager")
	}

	if !actorManager.Remove(&test2) {
		t.Errorf("incorrect recognition of Actors in ActorManager")
	}
}
