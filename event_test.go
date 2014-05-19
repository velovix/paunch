package paunch

import (
	"testing"
)

type testObject struct {
	collider Collider
	id       int
}

func (obj *testObject) GetColliders() []Collider {

	return []Collider{obj.collider}
}

func (obj *testObject) OnCollision(objectCollider, culpritCollider Collider, culprit interface{}) {

	obj.id = 1
}

func (obj *testObject) OnKeyboard(key Key, action Action) {

	if key == KeyUp && action == Press {
		obj.id = 2
	}
}

func (obj *testObject) OnMouseButton(button MouseButton, action Action, x, y int) {

	if button == MouseButtonMiddle && action == Press && x == 5 && y == 5 {
		obj.id = 3
	}
}

func (obj *testObject) OnMousePosition(x, y int) {

	if x == 5 && y == 5 {
		obj.id = 4
	}
}

func (obj *testObject) OnMouseEnterWindow(x, y int, entered bool) {

	if x == 0 && y == 0 && entered {
		obj.id = 5
	}
}

func (obj *testObject) OnWindowFocus(focused bool) {

	if focused {
		obj.id = 6
	}
}

func (obj *testObject) OnDraw() {

	obj.id = 7
}

func (obj *testObject) OnTick() {

	if obj.id == 0 {
		obj.id = 8
	}
}

func (obj *testObject) OnJoystickButton(button int, action Action) {

	if button == 1 && action == Press {
		obj.id = 9
	}
}

func (obj *testObject) OnJoystickAxis(device int, value float64) {

	if device == 0 && value == 0.5 {
		obj.id = 10
	}
}

func (obj *testObject) OnWindowResize(width, height int) {

	if width == 640 && height == 480 {
		obj.id = 11
	}
}

func TestEventManagerCollisions(t *testing.T) {

	eventManager := NewEventManager()

	test1 := testObject{NewPoint(0.0, 0.0), 0}
	test2 := testObject{NewPoint(0.0, 0.0), 0}
	test3 := testObject{NewPoint(1.0, 1.0), 0}

	eventManager.SetObjects([]interface{}{&test1, &test2, &test3})

	eventManager.RunCollisionEvent()

	if test1.id != 1 || test2.id != 1 {
		t.Errorf("method OnCollision not triggered")
	}

	if test3.id == 1 {
		t.Errorf("method OnCollision triggered incorrectly.")
	}
}

func TestEventManagerKeyboardInput(t *testing.T) {

	eventManager := NewEventManager()

	test := testObject{NewPoint(0.0, 0.0), 0}
	eventManager.SetObjects([]interface{}{&test})

	eventManager.RunKeyEvent(KeyUp, Press)

	if test.id != 2 {
		t.Errorf("method OnKeyboard not triggered")
	}
}

func TestEventManagerMouseInput(t *testing.T) {

	eventManager := NewEventManager()

	test := testObject{NewPoint(0.0, 0.0), 0}
	eventManager.SetObjects([]interface{}{&test})

	eventManager.RunMouseButtonEvent(MouseButtonMiddle, Press, 5, 5)

	if test.id != 3 {
		t.Errorf("method OnMouseButton not triggered")
	}
}

func TestEventManagerMousePosition(t *testing.T) {

	eventManager := NewEventManager()

	test := testObject{NewPoint(0.0, 0.0), 0}
	eventManager.SetObjects([]interface{}{&test})

	eventManager.RunMousePositionEvent(5, 5)

	if test.id != 4 {
		t.Errorf("method OnMousePosition not triggered")
	}
}

func TestObjectMouseEnterWindow(t *testing.T) {

	eventManager := NewEventManager()

	test := testObject{NewPoint(0.0, 0.0), 0}
	eventManager.SetObjects([]interface{}{&test})

	eventManager.RunMouseEnterWindowEvent(0, 0, true)

	if test.id != 5 {
		t.Errorf("method OnMouseEnterWindow not triggered")
	}
}

func TestObjectWindowFocused(t *testing.T) {

	eventManager := NewEventManager()

	test := testObject{NewPoint(0.0, 0.0), 0}
	eventManager.SetObjects([]interface{}{&test})

	eventManager.RunWindowFocusEvent(true)

	if test.id != 6 {
		t.Errorf("method OnWindowFocus not triggered")
	}
}

func TestObjectDraw(t *testing.T) {

	eventManager := NewEventManager()

	test := testObject{NewPoint(0.0, 0.0), 0}
	eventManager.SetObjects([]interface{}{&test})

	eventManager.RunDrawEvent()

	if test.id != 7 {
		t.Errorf("method OnDraw not triggered")
	}
}

func TestObjectTick(t *testing.T) {

	eventManager := NewEventManager()

	test := testObject{NewPoint(0.0, 0.0), 0}
	eventManager.SetObjects([]interface{}{&test})

	eventManager.RunTickEvent()

	if test.id != 8 {
		t.Errorf("method OnTick not triggered")
	}
}

func TestObjectJoystickButton(t *testing.T) {

	eventManager := NewEventManager()

	test := testObject{NewPoint(0.0, 0.0), 0}
	eventManager.SetObjects([]interface{}{&test})

	eventManager.RunJoystickButtonEvent(1, Press)

	if test.id != 9 {
		t.Errorf("method OnJoystickButton not triggered")
	}
}

func TestObjectJoystickAxis(t *testing.T) {

	eventManager := NewEventManager()

	test := testObject{NewPoint(0.0, 0.0), 0}
	eventManager.SetObjects([]interface{}{&test})

	eventManager.RunJoystickAxisEvent(0, 0.5)

	if test.id != 10 {
		t.Errorf("method OnJoystickAxis not triggered")
	}
}

func TestObjectWindowResize(t *testing.T) {

	eventManager := NewEventManager()

	test := testObject{NewPoint(0.0, 0.0), 0}
	eventManager.SetObjects([]interface{}{&test})

	eventManager.RunWindowResizeEvent(640, 480)

	if test.id != 11 {
		t.Errorf("method OnWindowResize not triggered")
	}
}
