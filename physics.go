package paunch

import (
	"math"
)

// Physics is an object meant to make the movement of multiple related movers,
// such as a Renderable and a Collision, easier. It also allows for easy
// management of multiple forces of movement at once.
type Physics struct {
	movers []Mover

	accel    Point
	friction Point

	forces map[string]Point
}

// NewPhysics creates a new Physics object.
func NewPhysics() Physics {

	var physics Physics
	physics.forces = make(map[string]Point)

	return physics
}

// AttachMover attaches a Mover to the Physics object. This Mover will be moved
// along with the rest of the Physics object as soon as it's attached.
func (physics *Physics) AttachMover(mover Mover) {

	physics.movers = append(physics.movers, mover)
}

// AddConstForce adds a constant force to the Physics object, which is taken
// into account every time the Calculate method is called. This might be used
// to simulate gravity or other such forces.
func (physics *Physics) AddConstForce(name string, force Point) {

	physics.forces[name] = force
}

// DeleteConstForce removes a constant force from the Physics object.
func (physics *Physics) DeleteConstForce(name string) {

	delete(physics.forces, name)
}

// Accelerate exerts a specified force upon the Physics object the next time
// the Calculate method is called.
func (physics *Physics) Accelerate(force Point) {

	physics.accel.X += force.X
	physics.accel.Y += force.Y
}

// SetFriction sets the friction value of the Physics object. Friction is a
// force that enfluences acceleration to move toward zero. This might be used
// to simulate the natural slowdown of an object rubbing against a surface.
func (physics *Physics) SetFriction(force Point) {

	physics.friction = force
}

// Calculate moves the Physics object given any specified constant forces,
// calls to the Accelerate method, and any leftover acceleration. Then,
// friction is applied to the resulting acceleration value.
func (physics *Physics) Calculate() {

	for _, val := range physics.forces {
		physics.accel.X += val.X
		physics.accel.Y += val.Y
	}

	for i := range physics.movers {
		physics.movers[i].Move(physics.accel.X, physics.accel.Y)
	}

	if math.Abs(physics.accel.X) >= math.Abs(physics.friction.X) {
		if physics.accel.X > 0 {
			physics.accel.X -= physics.friction.X
		} else {
			physics.accel.X += physics.friction.X
		}
	}
	if math.Abs(physics.accel.Y) >= math.Abs(physics.friction.Y) {
		if physics.accel.Y > 0 {
			physics.accel.Y -= physics.friction.Y
		} else {
			physics.accel.Y += physics.friction.Y
		}
	}
}
