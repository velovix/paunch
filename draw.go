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

// Clear clears the pixels on screen. This should probably be called before
// every new frame.
func Clear() error {
	gl.Clear(gl.COLOR_BUFFER_BIT)

	return checkForErrors()
}
