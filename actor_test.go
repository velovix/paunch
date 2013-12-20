package paunch

import (
	"testing"
)

type testActorObject struct {
	collider Collider
	id       int
}

func (obj *testActorObject) GetColliders() []Collider {

	return []Collider{obj.collider}
}

func (obj *testActorObject) OnCollision(objectCollider, culpritCollider Collider, culprit Actor) {

	obj.id = 1
}

func (obj *testActorObject) OnKeyboard(key, action int) {

	if key == KeyUp && action == KeyPress {
		obj.id = 2
	}
}

func (obj *testActorObject) Draw() {

}

func TestActorManager(t *testing.T) {

	var actorManager ActorManager

	test1 := testActorObject{NewPoint(0.0, 0.0), 0}
	test2 := testActorObject{NewPoint(0.0, 0.0), 0}

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

func TestActorManagerCollisions(t *testing.T) {

	var actorManager ActorManager

	test1 := testActorObject{NewPoint(0.0, 0.0), 0}
	test2 := testActorObject{NewPoint(0.0, 0.0), 0}
	test3 := testActorObject{NewPoint(1.0, 1.0), 0}

	actorManager.Add(&test1)
	actorManager.Add(&test2)
	actorManager.Add(&test3)

	actorManager.Run()

	if test1.id != 1 || test2.id != 1 {
		t.Errorf("method OnCollision not triggered")
	}

	if test3.id == 1 {
		t.Errorf("method OnCollision triggered incorrectly.")
	}
}

func TestActorManagerKeyboardInput(t *testing.T) {

	var actorManager ActorManager

	test := testActorObject{NewPoint(0.0, 0.0), 0}
	actorManager.Add(&test)

	actorManager.keyEvent(int(KeyUp), int(KeyPress))

	if test.id != 2 {
		t.Errorf("method OnKeyboard not triggered")
	}
}
