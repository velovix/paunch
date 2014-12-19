package paunch

// The Mover interface represents an object that can move and report it's
// positions.
type Mover interface {
	Move(x, y float64)
	SetPosition(x, y float64)
	Position() (x, y float64)
}
