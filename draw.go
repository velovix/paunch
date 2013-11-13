package paunch

import (
	"errors"
	"fmt"
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

// Renderable is an object that can be drawn on the screen
type Renderable struct {
	mode            int
	size            int
	vertex_buffer   gl.Uint
	texcoord_buffer gl.Uint
	texture_id      gl.Uint
}

func checkForErrors() error {

	var errList []gl.Enum
	for err := gl.GetError(); err != gl.NO_ERROR; {
		errList = append(errList, err)
	}

	if len(errList) == 0 {
		return nil
	} else {
		return errors.New(fmt.Sprintln("OpenGL Error(s): ", errList))
	}
}

// InitDraw sets up the drawing session for use.
func InitDraw(window Window) error {

	runtime.LockOSThread()

	if err := gl.Init(); err != nil {
		return errors.New("Error initializing OpenGL")
	}

	gl.ClearColor(0.0, 0.0, 0.0, 1.0)
	gl.Viewport(0, 0, gl.Sizei(window.Width), gl.Sizei(window.Height))

	return checkForErrors()
}

// NewRenderable returns a new Renderable object based on the specified shape
// type and verticies.
func NewRenderable(mode int, verticies []float32, texCoords []float32) (Renderable, error) {

	renderable := Renderable{mode, len(verticies), 0, 0, 0}

	gl.GenBuffers(1, &renderable.vertex_buffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, gl.Uint(renderable.vertex_buffer))
	gl.BufferData(gl.ARRAY_BUFFER, gl.Sizeiptr(len(verticies)*4), gl.Pointer(&verticies[0]), gl.STATIC_DRAW)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	if texCoords != nil {
		gl.GenBuffers(1, &renderable.texcoord_buffer)
		gl.BindBuffer(gl.ARRAY_BUFFER, gl.Uint(renderable.texcoord_buffer))
		gl.BufferData(gl.ARRAY_BUFFER, gl.Sizeiptr(len(texCoords)*4), gl.Pointer(&texCoords[0]), gl.STATIC_DRAW)
		gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	}

	return renderable, checkForErrors()
}

// DrawRenderable draws a Renderable
func DrawRenderable(renderable Renderable) {

	gl.BindBuffer(gl.ARRAY_BUFFER, renderable.vertex_buffer)
	gl.VertexAttribPointer(gl.Uint(0), 2, gl.FLOAT, gl.FALSE, 0, gl.Offset(nil, 0))
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	gl.EnableVertexAttribArray(gl.Uint(0))
	gl.DrawArrays(gl.Enum(renderable.mode), 0, gl.Sizei(renderable.size))
	gl.DisableVertexAttribArray(gl.Uint(0))
}

// Clear clears the pixels on screen. This should probably be called before
// every new frame.
func Clear() {
	gl.Clear(gl.COLOR_BUFFER_BIT)
}
