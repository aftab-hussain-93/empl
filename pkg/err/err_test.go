// Custom errors package to efficiently convert application errors into http errors/status code
package fault

import (
	"errors"
	"testing"
)

func TestFromError(t *testing.T) {
	t.Run("input is custom error", func(t *testing.T) {
		e := &Err{
			Code:          ErrInternalServer,
			Message:       "Internal Server Error",
			InternalError: errors.New("internal error"),
		}
		r, ok := FromError(e)
		if !ok {
			t.Error("error must be OK")
		}
		if r.Code != ErrInternalServer {
			t.Error("incorrect error code")
		}
		if r.InternalError.Error() != "internal error" {
			t.Error("incorrect internal error")
		}
	})
	t.Run("input is simple go error", func(t *testing.T) {
		e := errors.New("simple error")
		_, ok := FromError(e)
		if ok {
			t.Error("error must not be OK")
		}
	})
}
