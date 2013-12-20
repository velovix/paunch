package paunch

// Actor is an interface to be implemented by user-defined game objects.
// Used to make keeping track of multiple objects easier by automating the
// process of calling methods when common events happen.
type Actor interface {
	GetColliders() []Collider
	OnCollision(c1, c2 Collider, culprit Actor)
	OnKeyboard(key, action int)
	Draw()
}
