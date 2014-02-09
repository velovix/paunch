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

func (obj *testActorObject) OnKeyboard(key Key, action Action) {

	if key == KeyUp && action == Press {
		obj.id = 2
	}
}

func (obj *testActorObject) OnMouseButton(button MouseButton, action Action, x, y int) {

	if button == MouseButtonMiddle && action == Press && x == 5 && y == 5 {
		obj.id = 3
	}
}

func (obj *testActorObject) OnMousePosition(x, y int) {

	if x == 5 && y == 5 {
		obj.id = 4
	}
}

func (obj *testActorObject) OnMouseEnterWindow(x, y int, entered bool) {

	if x == 0 && y == 0 && entered {
		obj.id = 5
	}
}

func (obj *testActorObject) OnWindowFocus(focused bool) {

	if focused {
		obj.id = 6
	}
}

func (obj *testActorObject) OnDraw() {

	obj.id = 7
}

func (obj *testActorObject) OnTick() {

	if obj.id == 0 {
		obj.id = 8
	}
}

func (obj *testActorObject) OnJoystickButton(button int, action Action) {

	if button == 1 && action == Press {
		obj.id = 9
	}
}

func (obj *testActorObject) OnJoystickAxis(device int, value float64) {

	if device == 0 && value == 0.5 {
		obj.id = 10
	}
}

func TestActorManager(t *testing.T) {

	actorManager := NewActorManager()

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

	actorManager := NewActorManager()

	test1 := testActorObject{NewPoint(0.0, 0.0), 0}
	test2 := testActorObject{NewPoint(0.0, 0.0), 0}
	test3 := testActorObject{NewPoint(1.0, 1.0), 0}

	actorManager.Add(&test1)
	actorManager.Add(&test2)
	actorManager.Add(&test3)

	actorManager.RunCollisionEvent()

	if test1.id != 1 || test2.id != 1 {
		t.Errorf("method OnCollision not triggered")
	}

	if test3.id == 1 {
		t.Errorf("method OnCollision triggered incorrectly.")
	}
}

func TestActorManagerKeyboardInput(t *testing.T) {

	actorManager := NewActorManager()

	test := testActorObject{NewPoint(0.0, 0.0), 0}
	actorManager.Add(&test)

	actorManager.RunKeyEvent(KeyUp, Press)

	if test.id != 2 {
		t.Errorf("method OnKeyboard not triggered")
	}
}

func TestActorManagerMouseInput(t *testing.T) {

	actorManager := NewActorManager()

	test := testActorObject{NewPoint(0.0, 0.0), 0}
	actorManager.Add(&test)

	actorManager.RunMouseButtonEvent(MouseButtonMiddle, Press, 5, 5)

	if test.id != 3 {
		t.Errorf("method OnMouseButton not triggered")
	}
}

func TestActorManagerMousePosition(t *testing.T) {

	actorManager := NewActorManager()

	test := testActorObject{NewPoint(0.0, 0.0), 0}
	actorManager.Add(&test)

	actorManager.RunMousePositionEvent(5, 5)

	if test.id != 4 {
		t.Errorf("method OnMousePosition not triggered")
	}
}

func TestActorMouseEnterWindow(t *testing.T) {

	actorManager := NewActorManager()

	test := testActorObject{NewPoint(0.0, 0.0), 0}
	actorManager.Add(&test)

	actorManager.RunMouseEnterWindowEvent(0, 0, true)

	if test.id != 5 {
		t.Errorf("method OnMouseEnterWindow not triggered")
	}
}

func TestActorWindowFocused(t *testing.T) {

	actorManager := NewActorManager()

	test := testActorObject{NewPoint(0.0, 0.0), 0}
	actorManager.Add(&test)

	actorManager.RunWindowFocusEvent(true)

	if test.id != 6 {
		t.Errorf("method OnWindowFocus not triggered")
	}
}

func TestActorDraw(t *testing.T) {

	actorManager := NewActorManager()

	test := testActorObject{NewPoint(0.0, 0.0), 0}
	actorManager.Add(&test)

	actorManager.RunDrawEvent()

	if test.id != 7 {
		t.Errorf("method OnDraw not triggered")
	}
}

func TestActorTick(t *testing.T) {

	actorManager := NewActorManager()

	test := testActorObject{NewPoint(0.0, 0.0), 0}
	actorManager.Add(&test)

	actorManager.RunTickEvent()

	if test.id != 8 {
		t.Errorf("method OnTick not triggered")
	}
}

func TestActorJoystickButton(t *testing.T) {

	actorManager := NewActorManager()

	test := testActorObject{NewPoint(0.0, 0.0), 0}
	actorManager.Add(&test)

	actorManager.RunJoystickButtonEvent(1, Press)

	if test.id != 9 {
		t.Errorf("method OnJoystickButton not triggered")
	}
}

func TestActorJoystickAxis(t *testing.T) {

	actorManager := NewActorManager()

	test := testActorObject{NewPoint(0.0, 0.0), 0}
	actorManager.Add(&test)

	actorManager.RunJoystickAxisEvent(0, 0.5)

	if test.id != 10 {
		t.Errorf("method OnJoystickAxis not triggered")
	}
}
