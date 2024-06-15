package server

import (
	"encoding/json"
	"net/http/httptest"
	"testing"
)

func TestRespond(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		Respond(recorder, 429, map[string]any{"name": "John Doe", "id": 1234})

		if recorder.Code != 429 {
			t.Errorf("invalid response status, expected 429, got %v", recorder.Code)
		}
		resp := map[string]any{}
		if err := json.NewDecoder(recorder.Result().Body).Decode(&resp); err != nil {
			t.Error("invalid response body, cannot decode")
		}
		if resp["name"] != "John Doe" || resp["id"].(float64) != 1234 {
			t.Errorf("invalid response body, got invalid body %+v", resp)
		}
	})
}
