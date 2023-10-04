package customErrors

import (
	"net/http"

	"emperror.dev/errors"
)

func NewBadRequestError(message string) error {
	br := &badRequestError{
		CustomError: NewCustomError(nil, http.StatusBadRequest, message),
	}
	stackErr := errors.WithStackIf(br)

	return stackErr
}

func NewBadRequestErrorWrap(err error, message string) error {
	br := &badRequestError{
		CustomError: NewCustomError(err, http.StatusBadRequest, message),
	}
	stackErr := errors.WithStackIf(br)

	return stackErr
}

type badRequestError struct {
	CustomError
}

type BadRequestError interface {
	CustomError
}

func (b *badRequestError) isBadRequestError() bool {
	return true
}

func IsBadRequestError(err error) bool {
	var badRequestError *badRequestError
	// us, ok := grpc_errors.Cause(err).(BadRequestError)
	if errors.As(err, &badRequestError) {
		return badRequestError.isBadRequestError()
	}

	return false
}
