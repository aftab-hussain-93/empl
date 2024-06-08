package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/aftab-hussain-93/empl/types"
	"github.com/go-chi/chi/v5"
)

type handler struct {
	svc types.EmployeeService
}

func CreateHandler(svc types.EmployeeService) func(chi.Router) http.Handler {
	h := &handler{
		svc: svc,
	}
	return func(r chi.Router) http.Handler {
		r.Route("/employees", func(r chi.Router) {
			r.Post("/", h.create())
			r.Get("/", h.list())
			r.Get("/{id}", h.getByID())
			r.Delete("/{id}", h.delete())
		})
		return r
	}
}

func (h *handler) create() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		emp := &types.Employee{}
		err := bindJSON(r, emp)
		if err != nil {
			writeErrorResponse(w, err)
			return
		}
		emp, err = h.svc.CreateEmployee(r.Context(), emp)
		if err != nil {
			writeErrorResponse(w, err)
			return
		}
		writeResponse(w, http.StatusCreated, emp)
	}
}

func (h *handler) list() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (h *handler) getByID() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (h *handler) delete() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func writeErrorResponse(w http.ResponseWriter, err error) {
	status, resp := getErrorResponse(err)
	writeResponse(w, status, resp)
}

func getErrorResponse(err error) (int, any) {
	resp := map[string]string{
		"error": err.Error(),
	}

	return http.StatusInternalServerError, resp
}

func writeResponse(w http.ResponseWriter, statusCode int, body any) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

func bindJSON(r *http.Request, body any) error {
	return json.NewDecoder(r.Body).Decode(body)
}
