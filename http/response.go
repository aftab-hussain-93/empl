package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	fault "github.com/aftab-hussain-93/empl/err"
)

func appErrorToHTTPError(err error) (int, *fault.HTTPError) {
	// unhandled exception
	resp := &fault.HTTPError{
		Error: &fault.Err{
			Code:          fault.ErrInternalServer,
			Message:       "Internal Server Error",
			InternalError: err,
		},
	}
	if appErr, ok := fault.FromError(err); ok {
		resp := &fault.HTTPError{Error: appErr}
		return appErr.Code.GetHTTPStatusCode(), resp
	}

	return http.StatusInternalServerError, resp
}

func RespondWithError(w http.ResponseWriter, err error) {
	status, resp := appErrorToHTTPError(err)
	Respond(w, status, resp)
}

func Respond(w http.ResponseWriter, statusCode int, body any) {
	respBytes, err := json.Marshal(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, _ = w.Write(respBytes)
}
