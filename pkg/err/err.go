// Custom errors package to efficiently convert application errors into http errors/status code
package fault

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type HTTPError struct {
	Error *Err `json:"error"`
}

// Err is a custom error struct that implements the `error` interface
type Err struct {
	Code          AppErrorCode
	Message       string
	InternalError error
}

func (e *Err) MarshalJSON() ([]byte, error) {
	internalErr := ""
	if e.InternalError != nil {
		internalErr = e.InternalError.Error()
	}
	return json.Marshal(&struct {
		Code          string `json:"code"`
		Message       string `json:"message"`
		InternalError string `json:"internalError,omitempty"`
	}{
		Code:          string(e.Code),
		Message:       e.Message,
		InternalError: internalErr,
	})
}

// Error implements error.
func (e *Err) Error() string {
	m := fmt.Sprintf("code=%s|message=%s", e.Code, e.Message)
	if e.InternalError != nil {
		m += fmt.Sprintf("|internalError=%s", e.InternalError.Error())
	}
	return m
}

func New(code AppErrorCode, message string, internalError error) *Err {
	return &Err{code, message, internalError}
}

var _ error = (*Err)(nil)

// AppErrorCode represents an application error code
// every app error has a corresponding http status code
type AppErrorCode string

const (
	ErrBadRequest     AppErrorCode = "bad_request"
	ErrNotFound       AppErrorCode = "not_found"
	ErrInternalServer AppErrorCode = "internal"
)

func (ae AppErrorCode) GetHTTPStatusCode() int {
	switch ae {
	case ErrBadRequest:
		return http.StatusBadRequest
	case ErrNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}

// Checks if the error is an application error
// @Returns
// applicationError - application error if true
// isApplicationError - bool
func FromError(err error) (*Err, bool) {
	e := &Err{}
	if errors.As(err, &e) {
		return e, true
	}
	return nil, false
}
