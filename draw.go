package paunch

import (
	"errors"
	gl "github.com/chsc/gogl/gl33"
	"runtime"
)

func checkForErrors() error {

	err := OpenGLError{make([]gl.Enum, 0)}
	for errNumb := gl.GetError(); errNumb != gl.NO_ERROR; {
		err.ErrorCodes = append(err.ErrorCodes, errNumb)
		if len(err.ErrorCodes) >= 10 {
			err.ErrorCodes = append(err.ErrorCodes, 255)
			for errNumb := gl.GetError(); errNumb != gl.NO_ERROR; {
			}
			break
		}
	}

	if len(err.ErrorCodes) == 0 {
		return nil
	}

	return err
}

func initDraw() error {

	runtime.LockOSThread()

	if err := gl.Init(); err != nil {
		return errors.New("initializing OpenGL")
	}

	gl.ClearColor(0.0, 0.0, 0.0, 1.0)

	gl.Enable(gl.BLEND)
	gl.BlendFuncSeparate(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA, gl.ONE, gl.ZERO)

	return checkForErrors()
}

// Clear clears the pixels on screen. This should probably be called before
// every new frame.
func Clear() error {

	gl.Clear(gl.COLOR_BUFFER_BIT)

	return checkForErrors()
}

// SetClearColor sets the color the screen will become after a call to the
// Clear function.
func SetClearColor(r uint8, g uint8, b uint8) {

	gl.ClearColor(gl.Float(r)/255, gl.Float(g)/255, gl.Float(b)/255, 1.0)
}

// UseEffect sets the given Effect object for use in the following draw
// commands.
func UseEffect(effect *Effect) error {

	gl.UseProgram(effect.program)

	effect.uniforms = make(map[string]gl.Int)

	return checkForErrors()
}

// DisableEffects disables all effects.
func DisableEffects() {

	gl.UseProgram(0)
}
