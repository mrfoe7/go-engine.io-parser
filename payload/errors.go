package payload

import (
	"errors"
	"fmt"
)

var (
	//
	errPaused = retryError{"paused"}

	//
	errTimeout = errors.New("timeout")

	//
	errInvalidPayload = errors.New("invalid payload")

	//
	errOverlap = errors.New("overlap")
)

// Error is payload error.
type Error interface {
	Error() string
	Temporary() bool
}

// OpError is operation error.
type OpError struct {
	Op  string
	Err error
}

func newOpError(op string, err error) error {
	return &OpError{
		Op:  op,
		Err: err,
	}
}

func (e *OpError) Error() string {
	return fmt.Sprintf("%s: %s", e.Op, e.Err.Error())
}

// Temporary returns true if error can retry.
func (e *OpError) Temporary() bool {
	if oe, ok := e.Err.(Error); ok {
		return oe.Temporary()
	}
	return false
}

type retryError struct {
	err string
}

func (e retryError) Error() string {
	return e.err
}

func (e retryError) Temporary() bool {
	return true
}
