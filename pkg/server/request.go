package server

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator/v10"

	fault "github.com/aftab-hussain-93/empl/pkg/err"
)

var validate *validator.Validate

func init() {
	validate = validator.New(validator.WithRequiredStructEnabled())
}

// BindRequestBody - binds the request body to Go obj
// and validates the incoming request using the validator pkg
// @Params -
// r - http request
// obj - the object into which the request body must be deserialized; this should ideally be a pointer to the object
func BindRequestBody(body io.Reader, obj any) error {
	if body == nil {
		return nil
	}
	if obj == nil {
		return nil
	}
	if err := json.NewDecoder(body).Decode(obj); err != nil {
		return fault.New(fault.ErrBadRequest, "Bad Request! "+err.Error(), nil)
	}
	if err := validate.Struct(obj); err != nil {
		return fault.New(fault.ErrBadRequest, "Bad Request! "+err.Error(), nil)
	}
	return nil
}
