package paunch

import (
	"errors"
	"github.com/velovix/gl"
	"runtime"
)

var paunchGLVersion Version
var paunchEffect *Effect

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

func initDraw(version Version) error {

	runtime.LockOSThread()

	if err := gl.InitVersion10(); err != nil {
		return errors.New("initializing OpenGL 1.0")
	} else if err = gl.InitVersion11(); err != nil {
		return errors.New("initializing OpenGL 1.1")
	} else if err = gl.InitVersion12(); err != nil {
		return errors.New("initializing OpenGL 1.2")
	} else if err = gl.InitVersion13(); err != nil {
		return errors.New("initializing OpenGL 1.3")
	} else if err = gl.InitVersion14(); err != nil {
		return errors.New("initializing OpenGL 1.4")
	} else if err = gl.InitVersion15(); err != nil {
		return errors.New("initializing OpenGL 1.5")
	} else if err = gl.InitVersion20(); err != nil {
		return errors.New("initializing OpenGL 2.0")
	} else if err = gl.InitVersion21(); err != nil {
		return errors.New("initializing OpenGL 2.1")
	}

	paunchGLVersion = VersionOld

	if version == VersionNew || version == VersionAutomatic {
		if err := gl.InitVersion30(); err != nil {
			if version == VersionNew {
				return errors.New("initializing OpenGL 3.0")
			}
		} else {
			paunchGLVersion = VersionNew
		}
	}

	gl.ClearColor(0.0, 0.0, 0.0, 1.0)

	gl.Enable(gl.BLEND)
	gl.BlendFuncSeparate(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA, gl.ONE, gl.ZERO)

	return checkForErrors()
}

// GetVersion returns the current OpenGL version being used, which would either
// be VersionOld or VersionNew.
func GetVersion() Version {

	return paunchGLVersion
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

	paunchEffect = effect

	return checkForErrors()
}

// DisableEffects disables all effects.
func DisableEffects() {

	gl.UseProgram(0)
	paunchEffect = nil
}
