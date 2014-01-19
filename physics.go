package paunch

import (
	"math"
)

type physicsPoint struct {
	x, y float64
}

// Physics is an object meant to make the movement of multiple related movers,
// such as a Renderable and a Collision, easier. It also allows for easy
// management of multiple forces of movement at once.
type Physics struct {
	movers []Mover

	accel    physicsPoint
	maxAccel physicsPoint
	friction physicsPoint

	usingMaxAccel map[Axis]bool
	forces        map[string]physicsPoint
}

// NewPhysics creates a new Physics object.
func NewPhysics() Physics {

	var physics Physics
	physics.usingMaxAccel = make(map[Axis]bool)
	physics.forces = make(map[string]physicsPoint)

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
func (physics *Physics) AddConstForce(name string, forceX, forceY float64) {

	physics.forces[name] = physicsPoint{forceX, forceY}
}

// DeleteConstForce removes a constant force from the Physics object.
func (physics *Physics) DeleteConstForce(name string) {

	delete(physics.forces, name)
}

// Move moves all the members of the Physics object a specified distance.
func (physics *Physics) Move(x, y float64) {

	for _, val := range physics.movers {
		val.Move(x, y)
	}
}

// SetPosition sets the position of the Physics object. The starting point
// is the start point of each individual object.
func (physics *Physics) SetPosition(x, y float64) {

	for _, val := range physics.movers {
		val.SetPosition(x, y)
	}
}

// GetAcceleration returns the X and Y acceleration of the Physics object.
func (physics *Physics) GetAcceleration() (float64, float64) {

	return physics.accel.x, physics.accel.y
}

// Accelerate exerts a specified force upon the Physics object the next time
// the Calculate method is called.
func (physics *Physics) Accelerate(forceX, forceY float64) {

	physics.accel.x += forceX
	physics.accel.y += forceY
}

// SetAcceleration sets the acceleration of the Physics object on the specified
// axis.
func (physics *Physics) SetAcceleration(force float64, axis Axis) {

	switch axis {
	case X:
		physics.accel.x = force
	case Y:
		physics.accel.y = force
	}
}

// SetMaxAcceleration sets the maximum allowed acceleration of the Physics
// object on the specified axis. In situations where the object would normally
// go faster than the specified value, it will be set to the value instead.
func (physics *Physics) SetMaxAcceleration(force float64, axis Axis) {

	switch axis {
	case X:
		physics.maxAccel.x = force
	case Y:
		physics.maxAccel.y = force
	}

	physics.usingMaxAccel[axis] = true
}

// SetFriction sets the friction value of the Physics object. Friction is a
// force that enfluences acceleration to move toward zero. This might be used
// to simulate the natural slowdown of an object rubbing against a surface.
func (physics *Physics) SetFriction(forceX, forceY float64) {

	physics.friction = physicsPoint{forceX, forceY}
}

// Calculate moves the Physics object given any specified constant forces,
// calls to the Accelerate method, and any leftover acceleration. Then,
// friction is applied to the resulting acceleration value.
func (physics *Physics) Calculate() {

	for _, val := range physics.forces {
		physics.accel.x += val.x
		physics.accel.y += val.y
	}

	if math.Abs(physics.accel.x) > physics.maxAccel.x && physics.usingMaxAccel[X] {
		if physics.accel.x > 0 {
			physics.accel.x = physics.maxAccel.x
		} else {
			physics.accel.x = -physics.maxAccel.x
		}
	}
	if math.Abs(physics.accel.y) > physics.maxAccel.y && physics.usingMaxAccel[Y] {
		if physics.accel.y > 0 {
			physics.accel.y = physics.maxAccel.y
		} else {
			physics.accel.y = -physics.maxAccel.y
		}
	}

	for i := range physics.movers {
		physics.movers[i].Move(physics.accel.x, physics.accel.y)
	}

	if math.Abs(physics.accel.x) >= math.Abs(physics.friction.x) {
		if physics.accel.x > 0 {
			physics.accel.x -= physics.friction.x
		} else {
			physics.accel.x += physics.friction.x
		}
	} else {
		physics.accel.x = 0
	}

	if math.Abs(physics.accel.y) >= math.Abs(physics.friction.y) {
		if physics.accel.y > 0 {
			physics.accel.y -= physics.friction.y
		} else {
			physics.accel.y += physics.friction.y
		}
	} else {
		physics.accel.y = 0
	}
}
