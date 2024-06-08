package routes

import (
	"net/http"

	http_server "github.com/aftab-hussain-93/empl/http"
	"github.com/aftab-hussain-93/empl/service"
	"github.com/go-chi/chi/v5"
)

type handler struct {
	svc service.EmployeeService
}

func CreateHandler(svc service.EmployeeService) func(chi.Router) http.Handler {
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
		req := &service.EmployeeCreateRequest{}
		err := http_server.BindRequestBody(r, req)
		if err != nil {
			http_server.RespondWithError(w, err)
			return
		}
		resp, err := h.svc.CreateEmployee(r.Context(), req)
		if err != nil {
			http_server.RespondWithError(w, err)
			return
		}
		http_server.Respond(w, http.StatusCreated, resp)
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
