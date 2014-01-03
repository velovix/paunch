package paunch

import (
	"fmt"
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
