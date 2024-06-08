package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	fault "github.com/aftab-hussain-93/empl/err"
)

func appErrorToHTTPError(err error) (int, any) {
	// unhandled exception
	resp := map[string]map[string]string{
		"error": {
			"message":       "Internal Server Error",
			"internalError": err.Error(),
		},
	}
	if appErr, ok := fault.FromError(err); ok {
		resp := map[string]*fault.Err{
			"error": appErr,
		}
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
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	w.WriteHeader(statusCode)
	w.Write(respBytes)
}
