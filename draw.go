package paunch

import (
	"errors"
	"fmt"
	gl "github.com/chsc/gogl/gl33"
	"runtime"
)

// Constants defining different shapes that can be created with a Renderable.
const (
	Points        = gl.POINTS
	LineStrip     = gl.LINE_STRIP
	LineLoop      = gl.LINE_LOOP
	Lines         = gl.LINES
	TriangleStrip = gl.TRIANGLE_STRIP
	TriangleFan   = gl.TRIANGLE_FAN
	Triangles     = gl.TRIANGLES
)

func checkForErrors() error {

	var errList []gl.Enum
	for err := gl.GetError(); err != gl.NO_ERROR; {
		errList = append(errList, err)
	}

	if len(errList) == 0 {
		return nil
	}

	return errors.New(fmt.Sprintln("OpenGL Error(s): ", errList))
}

// InitDraw sets up the drawing session for use.
func InitDraw(window Window) error {

	runtime.LockOSThread()

	if err := gl.Init(); err != nil {
		return errors.New("initializing OpenGL")
	}

	gl.ClearColor(0.0, 0.0, 0.0, 1.0)
	gl.Viewport(0, 0, gl.Sizei(window.Width), gl.Sizei(window.Height))

	return checkForErrors()
}

// Clear clears the pixels on screen. This should probably be called before
// every new frame.
func Clear() error {
	gl.Clear(gl.COLOR_BUFFER_BIT)

	return checkForErrors()
}
