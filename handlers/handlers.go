package handlers

import (
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
			r.Post("", h.create())
			r.Get("", h.list())
			r.Get("{id}", h.getByID())
			r.Delete("{id}", h.delete())
		})
		return nil
	}
}

func (r *handler) create() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (r *handler) list() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (r *handler) getByID() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (r *handler) delete() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {}
}
