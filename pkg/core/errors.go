package core

import (
	"errors"
	"fmt"
)

type Error struct {
	Code    string
	Reason  string
	Message string
	Details string
}

const ErrInternal = "internal"
const ErrInvalidArgument = "invalid_argument"
const ErrUnauthenticated = "unauthenticated"
const ErrUnauthorized = "unauthorized"

func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Message, e.Details)
}

func NewErr(code string, reason string, msg string, err error) *Error {
	details := ""
	if err != nil {
		details = err.Error()
	}
	return &Error{
		Code:    code,
		Reason:  reason,
		Message: msg,
		Details: details,
	}
}

func FbErr(code string, reason string, msg string, err error) *Error {
	var cerr *Error
	ok := errors.As(err, &cerr)
	if ok {
		return cerr
	}
	return NewErr(code, reason, msg, err)
}
