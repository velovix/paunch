package paunch

// ActorCollider is an interface that requires methods that allow an
// EventManager to check the object's Collider objects against those of other
// objects. Objects that also implement this interface will automatically be
// checked when added to an EventManager.
type ActorCollider interface {
	GetColliders() []Collider
	OnCollision(c1, c2 Collider, culprit interface{})
}

// ActorKeyboarder is an interface that requires methods that allow an
// EventManager to call the OnKeyboard method of an object when a keyboard
// event happens. Objects that also implement this interface will automatically
// be called when appropriate after being added to an EventManager.
type ActorKeyboarder interface {
	OnKeyboard(key Key, action Action)
}

// ActorMouseButtoner is an interface that requires methods that allow an
// EventManager to call the OnMouseButton method of an object when a mouse
// button event happens. Objects that also implement this interface will
// automatically be called when appropriate after being added to
// an EventManager.
type ActorMouseButtoner interface {
	OnMouseButton(button MouseButton, action Action, x, y int)
}

// ActorMousePositioner is an interface that requires methods that allow an
// EventManager to call the OnMousePosition method of an object when a mouse
// position event happens. Objects that also implement this interface will
// automatically be called when appropriate after being added to an
// EventManager.
type ActorMousePositioner interface {
	OnMousePosition(x, y int)
}

// ActorMouseEnterWindower is an interface that requires methods that allow an
// EventManager to call on the OnMouseEnterWindow method of an object when a
// mouse enters or leaves a window. Objects that also implement this interface
// will automatically be called when appropriate after being added to an
// EventManager.
type ActorMouseEnterWindower interface {
	OnMouseEnterWindow(x, y int, entered bool)
}

// ActorWindowFocuser is an interface that requires methods that allow an
// EventManager to call on the OnWindowFocus method of an object when the user
// changes the focus of a window. Objects that also implement this interface
// will automatically be called when appropriate after being added to an
// EventManager.
type ActorWindowFocuser interface {
	OnWindowFocus(focused bool)
}

// ActorWindowResizer is an interface that requires methods that allow an
// EventManager to call on the OnWindowResize method of an object when the user
// changes the size of a window. Objects that also implement this interface
// will automatically be called when appropriate after being added to an
// EventManager.
type ActorWindowResizer interface {
	OnWindowResize(width, height int)
}

// ActorJoystickButtoner is an interface that requires methods that allow an
// EventManager to call on the OnJoystickButton method of an object when the
// user presses, holds, or releases a joystick button. Objects that also
// implement this interface will automatically be called when appropriate
// after being added to an EventManager.
type ActorJoystickButtoner interface {
	OnJoystickButton(button int, action Action)
}

// ActorJoystickAxiser is an interface that requires methods that allow an
// EventManager to call on the OnJoystickAxis method of an object when the
// user has at least on analog device on their joystick. Objects that also
// implement this interface will automatically be called when appropriate after
// being added to an EventManager.
type ActorJoystickAxiser interface {
	OnJoystickAxis(device int, value float64)
}

// ActorDrawer is an interface that requires methods that allow an EventManager
// to draw an object at every frame. Objects that also implement this interface
// will be autmatically called on every frame.
type ActorDrawer interface {
	OnDraw()
}

// ActorTicker is an interface that requires methods that allow an EventManager
// to call on the OnTick method with every tick of the EventManager. Objects
// that also implement this interface will automatically be called when
// appropriate after being added to an EventManager.
type ActorTicker interface {
	OnTick()
}
