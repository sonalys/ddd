package interfaces

import "fmt"

type httpError struct {
	Code   int    `json:"code"`
	Reason string `json:"reason"`
}

func newEncodeError(err error) httpError {
	return httpError{
		Code:   0,
		Reason: fmt.Sprintf("failed to encode response: %s", err),
	}
}

func newProcessError(err error) httpError {
	return httpError{
		Code:   1,
		Reason: fmt.Sprintf("failed to process request: %s", err),
	}
}

func newRequestError(err error) httpError {
	return httpError{
		Code:   2,
		Reason: fmt.Sprintf("invalid request: %s", err),
	}
}

func newNotFoundError(err error) httpError {
	return httpError{
		Code:   2,
		Reason: err.Error(),
	}
}
