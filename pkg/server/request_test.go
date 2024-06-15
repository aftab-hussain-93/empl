package server

import (
	"bytes"
	"testing"

	"github.com/aftab-hussain-93/empl/internal/service"
)

func TestBindRequestBody(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		bdy := []byte(`{  "name": "John Doe", "position": "trainee", "salary": 500 }`)
		resp := &service.EmployeeCreateRequest{}
		err := BindRequestBody(bytes.NewBuffer(bdy), resp)
		if err != nil {
			t.Error("bind request body failed, err", err.Error())
		}
		if resp.Name != "John Doe" || resp.Position != "trainee" || resp.Salary != 500 {
			t.Errorf("binding request body failed, got %+v", resp)
		}
	})
	t.Run("error", func(t *testing.T) {
		bdy := []byte(`{  "name": "John Doe", "position": "associate", "salary": 500 }`)
		resp := &service.EmployeeCreateRequest{}
		err := BindRequestBody(bytes.NewBuffer(bdy), resp)
		if err == nil {
			t.Error("test failed, expected bad request error")
		}
	})
}
