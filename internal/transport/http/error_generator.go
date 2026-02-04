package http

import (
	"errors"
	"net/http"
)

type Error struct {
	err    error
	Code   int
	Errors map[string]string
}

func (e Error) Error() string {
	if e.err != nil {
		return e.err.Error()
	}

	return ""
}

func NotFoundError(err error) Error {
	return Error{
		err:  err,
		Code: http.StatusNotFound,
	}
}

func InternalServerError(err error) Error {
	return Error{
		err:  err,
		Code: http.StatusInternalServerError,
	}
}

func BadRequestError(err error) Error {
	return Error{
		err:  err,
		Code: http.StatusBadRequest,
	}
}

func NotAuthenticatedError(err error) Error {
	return Error{
		err:  err,
		Code: http.StatusUnauthorized,
	}
}

func FailedValidationError(errorsBag map[string]string) Error {
	return Error{
		err:    errors.New("validation error"),
		Code:   422,
		Errors: errorsBag,
	}
}

func RateLimitExceeded() Error {
	err := errors.New("you have sent too many requests in short time, please try again later")

	return Error{
		err:  err,
		Code: http.StatusTooManyRequests,
	}
}

func BasicError(err error) Error {
	return InternalServerError(err)
}
