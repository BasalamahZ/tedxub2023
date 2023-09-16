package http

import (
	"errors"
)

// Followings are the known errors from Our Team HTTP handlers.
var (
	// errRequestTimeout is returned when processing time
	// has reached the timeout limit.
	errRequestTimeout = errors.New("REQUEST_TIMEOUT")
)
