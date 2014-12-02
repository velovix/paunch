package paunch

import (
	"math"
)

type physicsPoint struct {
	x, y float64
}

type force struct {
	magnitude physicsPoint
	active    bool
}

// Physics is an object meant to make the Movement of multiple related Movers,
// such as a Renderable and a Collision, easier. It also allows for easy
// management of multiple forces of Movement at once.
type Physics struct {
	Movers []Mover

	accel    physicsPoint
	maxAccel physicsPoint
	minAccel physicsPoint
	friction physicsPoint

	usingMaxAccel map[Axis]bool
	usingMinAccel map[Axis]bool
	forces        map[string]force
}

// NewPhysics creates a new Physics object.
func NewPhysics() *Physics {

	physics := &Physics{}
	physics.usingMaxAccel = make(map[Axis]bool)
	physics.usingMinAccel = make(map[Axis]bool)
	physics.forces = make(map[string]force)

	return physics
}

// AddForce adds a constant force to the Physics object, which is taken
// into account every time the Calculate method is called. The force is
// disabled by default.
func (physics *Physics) AddForce(name string, forceX, forceY float64) {

	physics.forces[name] = force{physicsPoint{forceX, forceY}, false}
}

// EnableForce makes the specified force active for future calls to the
// Calculate method.
func (physics *Physics) EnableForce(name string) {

	if _, ok := physics.forces[name]; ok {
		physics.forces[name] = force{physics.forces[name].magnitude, true}
	}
}

// DisableForce makes the specified force inactive for future calls to the
// Calculate method.
func (physics *Physics) DisableForce(name string) {

	if _, ok := physics.forces[name]; ok {
		physics.forces[name] = force{physics.forces[name].magnitude, false}
	}
}

// DeleteForce reMoves a constant force from the Physics object.
func (physics *Physics) DeleteForce(name string) {

	delete(physics.forces, name)
}

// Move Moves all the members of the Physics object a specified distance.
func (physics *Physics) Move(x, y float64) {

	for _, val := range physics.Movers {
		val.Move(x, y)
	}
}

// SetPosition sets the position of the Physics object relative to the Physics
// object's first Mover.
func (physics *Physics) SetPosition(x, y float64) {

	if len(physics.Movers) == 0 {
		return
	}

	xDisp, yDisp := physics.Movers[0].GetPosition()
	xDisp = x - xDisp
	yDisp = y - yDisp

	for _, val := range physics.Movers {
		val.Move(xDisp, yDisp)
	}
}

// GetPosition returns the position of the Physics object's first Mover.
func (physics *Physics) GetPosition() (x, y float64) {

	if len(physics.Movers) == 0 {
		return 0, 0
	}

	return physics.Movers[0].GetPosition()
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

// SetMinAcceleration sets the minimum allowed acceleration of the Physics
// object on the specified axis. In situations where the object would normally
// go slower than the specified value, it will be set to the value instead.
func (physics *Physics) SetMinAcceleration(force float64, axis Axis) {

	switch axis {
	case X:
		physics.minAccel.x = force
	case Y:
		physics.minAccel.y = force
	}

	physics.usingMinAccel[axis] = true
}

// SetFriction sets the friction value of the Physics object. Friction is a
// force that enfluences acceleration to Move toward zero. This might be used
// to simulate the natural slowdown of an object rubbing against a surface.
func (physics *Physics) SetFriction(forceX, forceY float64) {

	physics.friction = physicsPoint{forceX, forceY}
}

// Calculate Moves the Physics object given any specified constant forces,
// calls to the Accelerate method, and any leftover acceleration. Then,
// friction is applied to the resulting acceleration value.
func (physics *Physics) Calculate() {

	for _, val := range physics.forces {
		if val.active {
			physics.accel.x += val.magnitude.x
			physics.accel.y += val.magnitude.y
		}
	}

	if physics.accel.x > physics.maxAccel.x && physics.usingMaxAccel[X] {
		physics.accel.x = physics.maxAccel.x
	} else if physics.accel.x < physics.minAccel.x && physics.usingMinAccel[X] {
		physics.accel.x = physics.minAccel.x
	}

	if physics.accel.y > physics.maxAccel.y && physics.usingMaxAccel[Y] {
		physics.accel.y = physics.maxAccel.y
	} else if physics.accel.y < physics.minAccel.y && physics.usingMinAccel[Y] {
		physics.accel.y = physics.minAccel.y
	}

	for i := range physics.Movers {
		physics.Movers[i].Move(physics.accel.x, physics.accel.y)
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
