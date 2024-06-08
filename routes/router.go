package routes

import (
	"net/http"
	"strconv"

	fault "github.com/aftab-hussain-93/empl/err"
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
	type Response struct {
		Employees  []*service.Employee `json:"employees"`
		Count      int                 `json:"count"`
		PageNumber int                 `json:"page"`
		Total      int64               `json:"total"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		pageSize := 0
		if pageSizeStr := r.URL.Query().Get("size"); pageSizeStr != "" {
			size, err := strconv.Atoi(pageSizeStr)
			if err != nil {
				err := fault.New(fault.ErrBadRequest, "Bad Request! Invalid query param `size`, must be an positive integer", err)
				http_server.RespondWithError(w, err)
				return
			}
			if size < 0 {
				size = 0
			}
			pageSize = size
		}
		page := 1
		if pageNumberStr := r.URL.Query().Get("page"); pageNumberStr != "" {
			num, err := strconv.Atoi(pageNumberStr)
			if err != nil {
				err := fault.New(fault.ErrBadRequest, "Bad Request! Invalid query param `page`, must be an positive integer", err)
				http_server.RespondWithError(w, err)
				return
			}
			if num <= 0 {
				num = 1
			}
			page = num
		}

		emps, totalEmps, err := h.svc.ListEmployees(r.Context(), int64(pageSize), int64(page))
		if err != nil {
			http_server.RespondWithError(w, err)
			return
		}

		resp := Response{
			Employees:  emps,
			Total:      totalEmps,
			Count:      len(emps),
			PageNumber: page,
		}

		http_server.Respond(w, http.StatusOK, resp)
	}
}

func (h *handler) getByID() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (h *handler) delete() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {}
}
