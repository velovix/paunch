package paunch

import (
	"fmt"
	gl "github.com/chsc/gogl/gl33"
)

type LineGetPointFromErrorType int

const (
	_ LineGetPointFromErrorType = iota
	OutsideLineRangeError
	UndefinedSlopeError
)

type LineGetPointFromError struct {
	Number float64
	Line   *Line
	Type   LineGetPointFromErrorType
}

func (err LineGetPointFromError) Error() string {

	switch err.Type {
	case OutsideLineRangeError:
		return fmt.Sprintf("value %f is outside line range", err.Number)
	case UndefinedSlopeError:
		return "no valid Point found on Line with undefined slope"
	default:
		return "undefined error"
	}
}

type OpenGLError struct {
	ErrorCodes []gl.Enum
}

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
		default:
			message += fmt.Sprintf("Unknown error %i ", val)
		}
	}

	return message
}
