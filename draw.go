package paunch

import (
	"errors"
	gl "github.com/chsc/gogl/gl33"
	"runtime"
)

const (
	POINTS         = gl.POINTS
	LINE_STRIP     = gl.LINE_STRIP
	LINE_LOOP      = gl.LINE_LOOP
	LINES          = gl.LINES
	TRIANGLE_STRIP = gl.TRIANGLE_STRIP
	TRIANGLE_FAN   = gl.TRIANGLE_FAN
	TRIANGLES      = gl.TRIANGLES
)

type Renderable struct {
	mode   int
	size   int
	buffer gl.Uint
}

type Draw struct {
	va_shape int
}

func (draw *Draw) Init(window Window) error {

	runtime.LockOSThread()

	if err := gl.Init(); err != nil {
		return errors.New("Error initializing OpenGL")
	}

	gl.ClearColor(0.0, 0.0, 0.0, 1.0)
	gl.Viewport(0, 0, gl.Sizei(window.Width), gl.Sizei(window.Height))

	draw.va_shape = 0

	return nil
}

func (draw *Draw) NewRenderable(mode int, verticies []float32) Renderable {

	renderable := Renderable{mode, len(verticies), 0}

	gl.GenBuffers(1, &renderable.buffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, gl.Uint(renderable.buffer))
	gl.BufferData(gl.ARRAY_BUFFER, gl.Sizeiptr(len(verticies)*4), gl.Pointer(&verticies[0]), gl.STATIC_DRAW)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	return renderable
}

func DrawVerticies(mode int, verticies []float32) {

	var buffer_id gl.Uint
	gl.GenBuffers(1, &buffer_id)
	gl.BindBuffer(gl.ARRAY_BUFFER, buffer_id)
	gl.BufferData(gl.ARRAY_BUFFER, gl.Sizeiptr(len(verticies)*8), gl.Pointer(&verticies[0]), gl.STATIC_DRAW)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, gl.FALSE, 0, gl.Offset(nil, 0))
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	gl.EnableVertexAttribArray(0)

	gl.DrawArrays(gl.Enum(mode), 0, gl.Sizei(len(verticies)))

	gl.DisableVertexAttribArray(0)
}

func (draw *Draw) DrawRenderable(renderable Renderable) {

	gl.BindBuffer(gl.ARRAY_BUFFER, renderable.buffer)
	gl.VertexAttribPointer(gl.Uint(draw.va_shape), 3, gl.FLOAT, gl.FALSE, 0, gl.Offset(nil, 0))
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	gl.EnableVertexAttribArray(gl.Uint(draw.va_shape))
	gl.DrawArrays(gl.Enum(renderable.mode), 0, gl.Sizei(renderable.size))
	gl.DisableVertexAttribArray(gl.Uint(draw.va_shape))
}

func Clear() {
	gl.Clear(gl.COLOR_BUFFER_BIT)
}
