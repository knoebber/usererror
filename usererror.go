package usererror

import (
	"errors"
	"fmt"
)

// Error returns a message to the user.
type Error struct {
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

// New returns a new usererror.
func New(message string) *Error {
	return &Error{Message: message}
}

// Format returns a formatted usererror.
func Format(format string, a ...interface{}) *Error {
	return New(fmt.Sprintf(format, a...))
}

// Convert attempts to convert err to a usererror.
func Convert(err error) (result *Error) {
	errors.As(err, &result)
	return
}
