package paunch

// The Mover interface includes all structs that implement the Move and
// SetPosition methods.
type Mover interface {
	Move(x, y float64)
	SetPosition(x, y float64)
	GetPosition() (x, y float64)
}
