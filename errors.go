package paunch

import (
	"fmt"
	gl "github.com/chsc/gogl/gl21"
)

type lineGetPointFromErrorType int

const (
	_ lineGetPointFromErrorType = iota
	outsideLineRangeError
	undefinedSlopeError
)

type lineGetPointFromError struct {
	Number float64
	line   *line
	Type   lineGetPointFromErrorType
}

func (err lineGetPointFromError) Error() string {

	switch err.Type {
	case outsideLineRangeError:
		return fmt.Sprintf("value %f is outside line range", err.Number)
	case undefinedSlopeError:
		return "no valid point found on line with undefined slope"
	default:
		return fmt.Sprintf("unknown error %v", err.Type)
	}
}

// OpenGLError is an error-implementing object describing sets of OpenGL
// errors.
type OpenGLError struct {
	ErrorCodes []gl.Enum
}

// Error returns a descriptive string.
func (err OpenGLError) Error() string {

	message := "openGL errors: "

	for _, val := range err.ErrorCodes {
		switch val {
		case gl.INVALID_ENUM:
			message += "INVALID_ENUM "
		case gl.INVALID_VALUE:
			message += "INVALID_VALUE "
		case gl.INVALID_OPERATION:
			message += "INVALID_OPERATION "
		case gl.STACK_OVERFLOW:
			message += "STACK_OVERFLOW "
		case gl.STACK_UNDERFLOW:
			message += "STACK_UNDERFLOW "
		case gl.OUT_OF_MEMORY:
			message += "OUT_OF_MEMORY "
		case 255:
			message += "..."
		default:
			message += fmt.Sprintf("Unknown error %v ", val)
		}
	}

	return message
}
