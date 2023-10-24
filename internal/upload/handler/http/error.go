package http

import (
	"errors"
)

// Followings are the known errors from upload HTTP handlers.
var (
	// errBadRequest is returned when the given request is
	// bad/invalid.
	errBadRequest = errors.New("BAD_REQUEST")

	// errInternalServerError is returned when there is an
	// internal error.
	errInternalServerError = errors.New("INTERNAL_SERVER_ERROR")

	// errFileTooLarge is returned when the request giving
	// file that size are bigger than max file size.
	errFileTooLarge = errors.New("FILE_TOO_LARGE")

	// errMethodNotAllowed is returned when accessing not
	// allowed HTTP method.
	errMethodNotAllowed = errors.New("METHOD_NOT_ALLOWED")

	// errRequestTimeout is returned when processing time has
	// reached the timeout limit.
	errRequestTimeout = errors.New("REQUEST_TIMEOUT")
)

var (
	// mapHTTPError maps service error into HTTP error that
	// categorize as bad request error.

	// Internal server error-related should not be mapped here,
	// and the handler should just return `errInternal` as the
	// error instead
	mapHTTPError = map[error]error{}
)
