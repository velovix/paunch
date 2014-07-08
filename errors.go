package paunch

import (
	"fmt"
	"github.com/velovix/gl"
)

// LineGetPointFromErrorType are the types of errors a LineGetPointFromError
// can represent.
type LineGetPointFromErrorType int

// LineGetPointFromError error type IDs.
const (
	_ LineGetPointFromErrorType = iota
	OutsideLineRangeError
	UndefinedSlopeError
)

// LineGetPointFromError is an error-implementing object returned by methods
// GetPointFromX and GetPointFromY.
type LineGetPointFromError struct {
	Number float64
	Line   *Line
	Type   LineGetPointFromErrorType
}

// Error returns a descriptive string.
func (err LineGetPointFromError) Error() string {

	switch err.Type {
	case OutsideLineRangeError:
		return fmt.Sprintf("value %f is outside line range", err.Number)
	case UndefinedSlopeError:
		return "no valid Point found on Line with undefined slope"
	default:
		return fmt.Sprintf("unknown error %i", err.Type)
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
		case gl.INVALID_FRAMEBUFFER_OPERATION:
			message += "INVALID_FRAMEBUFFER_OPERATION "
		case 255:
			message += "..."
		default:
			message += fmt.Sprintf("Unknown error %i ", val)
		}
	}

	return message
}
