package paunch

// Actor is an interface to be implemented by user-defined game objects.
// Used to make keeping track of multiple objects easier by automating the
// process of calling methods when common events happen.
type Actor interface{}

// ActorCollider is an interface that requires methods that allow an
// ActorManager to check the object's Collider objects against those of other
// Actors. Actor-implementing objects that also implement this interface will
// automatically be checked when added to an ActorManager.
type ActorCollider interface {
	GetColliders() []Collider
	OnCollision(c1, c2 Collider, culprit Actor)
}

// ActorKeyboarder is an interface that requires methods that allow an
// ActorManager to call the OnKeyboard method of an Actor when a keyboard event
// happens. Actor-implementing objects that also implement this interface will
// automatically be called when appropriate after being added to an
// ActorManager.
type ActorKeyboarder interface {
	OnKeyboard(key, action int)
}

// ActorMouseButtoner is an interface that requires methods that allow an
// ActorManager to call the OnMouseButton method of an Actor when a mouse
// button event happens. Actor-implementing objects that also implement this
// interface will automatically be called when appropriate after being added to
// an ActorManager.
type ActorMouseButtoner interface {
	OnMouseButton(button, action, x, y int)
}

// ActorMousePositioner is an interface that requires methods that allow an
// ActorManager to call the OnMousePosition method of an Actor when a mouse
// position event happens. Actor-implementing objects that also implement this
// interface will automatically be called when appropriate after being added to
// an ActorManager.
type ActorMousePositioner interface {
	OnMousePosition(x, y int)
}

// ActorMouseEnterWindower is an interface that requires methods that allow an
// ActorManager to call on the OnMouseEnterWindow method of an Actor when a
// mouse enters or leaves a window. Actor-implementing objects that also
// implement this interface will automatically be called when appropriate after
// being added to an ActorManager.
type ActorMouseEnterWindower interface {
	OnMouseEnterWindow(x, y int, entered bool)
}

// ActorWindowFocuser in an interface that requires methods that allow an
// ActorManager to call on the OnWindowFocus method of an Actor when the user
// changes the focus of a window. Actor-implementing objects that also
// implement this interface will automatically be called when appropriate after
// being added to an ActorManager.
type ActorWindowFocuser interface {
	OnWindowFocus(focused bool)
}

// ActorDrawer is an interface that requires methods that allow an ActorManager
// to draw an actor at every frame. Actor-implementing objects that also
// implement this interface will be autmatically called on every frame.
type ActorDrawer interface {
	OnDraw()
}

// ActorTicker is an interface that requires methods that allow an ActorManager
// to call on the OnTick method with every tick of the ActorManager.
// Actor-implementing objects that also implement this interface will
// automatically be called when appropriate after being added to an
// ActorManager.
type ActorTicker interface {
	OnTick()
}
