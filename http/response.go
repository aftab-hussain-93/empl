package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	fault "github.com/aftab-hussain-93/empl/err"
)

func appErrorToHTTPError(err error) (int, *fault.HTTPError) {
	// check for application errors
	if appErr, ok := fault.FromError(err); ok {
		resp := &fault.HTTPError{Error: appErr}
		return appErr.Code.GetHTTPStatusCode(), resp
	}

	// unhandled exception or non-application errors
	resp := &fault.HTTPError{
		Error: &fault.Err{
			Code:          fault.ErrInternalServer,
			Message:       "Internal Server Error",
			InternalError: err,
		},
	}
	return http.StatusInternalServerError, resp
}

// RespondWithError - parses the error and writes to response in a stand error response format
func RespondWithError(w http.ResponseWriter, err error) {
	status, resp := appErrorToHTTPError(err)
	Respond(w, status, resp)
}

func Respond(w http.ResponseWriter, statusCode int, body any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if body != nil {
		respBytes, err := json.Marshal(body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
			return
		}

		_, _ = w.Write(respBytes)
	}
}
