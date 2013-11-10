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

type DrawFloat gl.Float

func InitDraw(window Window) error {

	runtime.LockOSThread()

	if err := gl.Init(); err != nil {
		return errors.New("Error initializing OpenGL")
	}

	gl.ClearColor(0.0, 0.0, 0.0, 1.0)
	gl.Viewport(0, 0, gl.Sizei(window.Width), gl.Sizei(window.Height))

	return nil
}

func Draw(mode int, verts []DrawFloat) {

	var buffer_id gl.Uint
	gl.GenBuffers(1, &buffer_id)
	gl.BindBuffer(gl.ARRAY_BUFFER, buffer_id)
	gl.BufferData(gl.ARRAY_BUFFER, gl.Sizeiptr(len(verts)*8), gl.Pointer(&verts[0]), gl.STATIC_DRAW)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	gl.VertexAttribPointer(0, 3, gl.FLOAT, gl.FALSE, 0, gl.Offset(gl.Pointer(&verts[0]), 0))
	gl.EnableVertexAttribArray(0)

	gl.DrawArrays(gl.Enum(mode), 0, gl.Sizei(len(verts)))

	gl.DisableVertexAttribArray(0)
}
