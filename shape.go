package paunch

import (
	gl "github.com/chsc/gogl/gl21"
)

// Shape is an object that represents a vector shape, such as a triangle or
// other polygon, that can be drawn on screen.
type Shape struct {
	mode         gl.Enum
	size         int
	vertexBuffer gl.Uint
	verticies    []float32
}

// NewShape creates a new Shape object based on the verticies and shape type.
func NewShape(shapeType ShapeType, verticies []float64) (*Shape, error) {

	verticies32 := make([]float32, len(verticies))
	for i, val := range verticies {
		verticies32[i] = float32(val)
	}

	shape := &Shape{mode: gl.Enum(shapeType), size: len(verticies), vertexBuffer: 0, verticies: verticies32}

	gl.GenBuffers(1, &shape.vertexBuffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, gl.Uint(shape.vertexBuffer))
	gl.BufferData(gl.ARRAY_BUFFER, gl.Sizeiptr(shape.size*4), gl.Pointer(&shape.verticies[0]), gl.DYNAMIC_DRAW)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	return shape, checkForErrors()
}

// NewShapeFromShape creates a copy of an existing Shape object.
func NewShapeFromShape(copyShape *Shape) (*Shape, error) {

	shape := &Shape{mode: copyShape.mode, size: copyShape.size, verticies: make([]float32, len(copyShape.verticies))}

	copy(shape.verticies, copyShape.verticies)

	gl.GenBuffers(1, &shape.vertexBuffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, shape.vertexBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, gl.Sizeiptr(len(shape.verticies)*4), gl.Pointer(&shape.verticies[0]), gl.DYNAMIC_DRAW)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	return shape, checkForErrors()
}

// SetScaling sets the scaling factor of the Shape object. For instance, an x
// and y scale value of two will make the Shape object twice as large.
func (shape *Shape) SetScaling(xScale, yScale float64) {

	verticies := make([]float32, len(shape.verticies))

	xTransform := shape.verticies[0] - (shape.verticies[0] * float32(xScale))
	yTransform := shape.verticies[1] - (shape.verticies[1] * float32(yScale))

	for i := range verticies {
		if i%2 == 0 {
			verticies[i] = shape.verticies[i] * float32(xScale)
			verticies[i] += xTransform
		} else {
			verticies[i] = shape.verticies[i] * float32(yScale)
			verticies[i] += yTransform
		}
	}

	gl.BindBuffer(gl.ARRAY_BUFFER, shape.vertexBuffer)
	gl.BufferSubData(gl.ARRAY_BUFFER, 0, gl.Sizeiptr(len(verticies)*4), gl.Pointer(&verticies[0]))
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
}

// Draw draws the Shape object.
func (shape *Shape) Draw() error {

	gl.BindBuffer(gl.ARRAY_BUFFER, shape.vertexBuffer)
	vertexAttribLoc := gl.GetAttribLocation(paunchEffect.program, gl.GLString("position"))
	gl.VertexAttribPointer(gl.Uint(vertexAttribLoc), 2, gl.FLOAT, gl.FALSE, 0, gl.Offset(nil, 0))
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	gl.EnableVertexAttribArray(gl.Uint(vertexAttribLoc))
	gl.DrawArrays(shape.mode, 0, gl.Sizei(shape.size/2))
	gl.DisableVertexAttribArray(gl.Uint(vertexAttribLoc))
	//gl.DisableVertexAttribArray(gl.Uint(1))

	//gl.BindTexture(gl.TEXTURE_2D, 0)

	return checkForErrors()
}

// Move moves the Shape object a specified distance.
func (shape *Shape) Move(x, y float64) {

	for i := 0; i < len(shape.verticies); i += 2 {
		shape.verticies[i] += float32(x)
		shape.verticies[i+1] += float32(y)
	}

	gl.BindBuffer(gl.ARRAY_BUFFER, shape.vertexBuffer)
	gl.BufferSubData(gl.ARRAY_BUFFER, 0, gl.Sizeiptr(len(shape.verticies)*4), gl.Pointer(&shape.verticies[0]))
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
}

// SetPosition sets the position of the Shape object relative to first
// specified vertex.
func (shape *Shape) SetPosition(x, y float64) {

	xDisp := x - float64(shape.verticies[0])
	yDisp := y - float64(shape.verticies[1])

	shape.Move(xDisp, yDisp)
}

// GetPosition returns the X and Y position relative to the first specified
// vertex.
func (shape *Shape) GetPosition() (x, y float64) {

	return float64(shape.verticies[0]), float64(shape.verticies[1])
}
