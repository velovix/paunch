package paunch

// CollisionEventResponder is an interface that requires methods that allow an
// EventManager to check the object's Collider objects against those of other
// objects. Objects that implement this interface will automatically be checked
// when added to an EventManager.
type CollisionEventResponder interface {
	GetColliders() []Collider
	OnCollision(c1, c2 Collider, culprit interface{})
}

// KeyboardEventResponder is an interface that requires methods that allow an
// EventManager to call the OnKeyboard method of an object when a keyboard
// event happens. Objects that implement this interface will automatically be
// called when appropriate after being added to an EventManager.
type KeyboardEventResponder interface {
	OnKeyboard(key Key, action Action)
}

// MouseButtonEventResponder is an interface that requires methods that allow
// an EventManager to call the OnMouseButton method of an object when a mouse
// button event happens. Objects that implement this interface will
// automatically be called when appropriate after being added to an
// EventManager.
type MouseButtonEventResponder interface {
	OnMouseButton(button MouseButton, action Action, x, y int)
}

// MousePositionEventResponder is an interface that requires methods that allow
// an EventManager to call the OnMousePosition method of an object when a mouse
// position event happens. Objects that implement this interface will
// automatically be called when appropriate after being added to an
// EventManager.
type MousePositionEventResponder interface {
	OnMousePosition(x, y int)
}

// MouseEnterWindowResponder is an interface that requires methods that allow
// an EventManager to call on the OnMouseEnterWindow method of an object when a
// mouse enters or leaves a window. Objects that implement this interface will
// automatically be called when appropriate after being added to an
// EventManager.
type MouseEnterWindowResponder interface {
	OnMouseEnterWindow(x, y int, entered bool)
}

// WindowFocusEventResponder is an interface that requires methods that allow
// an EventManager to call on the OnWindowFocus method of an object when the
// user changes the focus of a window. Objects that implement this interface
// will automatically be called when appropriate after being added to an
// EventManager.
type WindowFocusEventResponder interface {
	OnWindowFocus(focused bool)
}

// WindowResizeEventResponder is an interface that requires methods that allow
// an EventManager to call on the OnWindowResize method of an object when the
// user changes the size of a window. Objects that implement this interface
// will automatically be called when appropriate after being added to an
// EventManager.
type WindowResizeEventResponder interface {
	OnWindowResize(width, height int)
}

// JoystickButtonEventResponder is an interface that requires methods that
// allow an EventManager to call on the OnJoystickButton method of an object
// when the user presses, holds, or releases a joystick button. Objects that
// implement this interface will automatically be called when appropriate after
// being added to an EventManager.
type JoystickButtonEventResponder interface {
	OnJoystickButton(button int, action Action)
}

// JoystickAxisEventResponder is an interface that requires methods that allow
// an EventManager to call on the OnJoystickAxis method of an object when the
// user has at least on analog device on their joystick. Objects that implement
// this interface will automatically be called when appropriate after being
// added to an EventManager.
type JoystickAxisEventResponder interface {
	OnJoystickAxis(device int, value float64)
}

// CharacterEventResponder is an interface that requires methods that allow an
// EventManager to call on the OnCharacter method of an object when the user
// inputs a valid unicode character. Objects that implement this interface will
// automatically be called when appropriate after being added to an
// EventManager.
type CharacterEventResponder interface {
	OnCharacter(character rune)
}

// DrawEventResponder is an interface that requires methods that allow an
// EventManager to draw an object at every frame. Objects that implement this
// interface will be autmatically called on every frame.
type DrawEventResponder interface {
	OnDraw()
}

// TickEventResponder is an interface that requires methods that allow an
// EventManager to call on the OnTick method with every tick of the
// EventManager. Objects that implement this interface will automatically be
// called when appropriate after being added to an EventManager.
type TickEventResponder interface {
	OnTick()
}
