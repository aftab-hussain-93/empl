package routes

import (
	"net/http"
	"strconv"

	"github.com/aftab-hussain-93/empl/internal/service"
	fault "github.com/aftab-hussain-93/empl/pkg/err"
	server "github.com/aftab-hussain-93/empl/pkg/server"
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
			r.Put("/{id}", h.update())
			r.Delete("/{id}", h.delete())
		})
		r.Get("/status", func(w http.ResponseWriter, r *http.Request) {
			server.Respond(w, http.StatusOK, map[string]string{"status": "ok"})
		})
		return r
	}
}

func (h *handler) create() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &service.EmployeeCreateRequest{}
		err := server.BindRequestBody(r.Body, req)
		if err != nil {
			server.RespondWithError(w, err)
			return
		}
		resp, err := h.svc.CreateEmployee(r.Context(), req)
		if err != nil {
			server.RespondWithError(w, err)
			return
		}
		server.Respond(w, http.StatusCreated, resp)
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
				server.RespondWithError(w, err)
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
				server.RespondWithError(w, err)
				return
			}
			if num <= 0 {
				num = 1
			}
			page = num
		}

		emps, totalEmps, err := h.svc.ListEmployees(r.Context(), int64(pageSize), int64(page))
		if err != nil {
			server.RespondWithError(w, err)
			return
		}

		resp := Response{
			Employees:  emps,
			Total:      totalEmps,
			Count:      len(emps),
			PageNumber: page,
		}

		server.Respond(w, http.StatusOK, resp)
	}
}

func (h *handler) getByID() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" {
			server.RespondWithError(w, fault.New(fault.ErrBadRequest, "id not provided", nil))
			return
		}
		empID, err := strconv.Atoi(id)
		if err != nil {
			server.RespondWithError(w, fault.New(fault.ErrBadRequest, "invalid id provided", nil))
			return
		}
		emp, err := h.svc.GetEmployeeByID(r.Context(), empID)
		if err != nil {
			server.RespondWithError(w, err)
			return
		}
		server.Respond(w, http.StatusOK, emp)
	}
}

func (h *handler) update() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" {
			server.RespondWithError(w, fault.New(fault.ErrBadRequest, "id not provided", nil))
			return
		}
		empID, err := strconv.Atoi(id)
		if err != nil {
			server.RespondWithError(w, fault.New(fault.ErrBadRequest, "invalid id provided", nil))
			return
		}
		req := &service.EmployeeUpdateRequest{}
		if err := server.BindRequestBody(r.Body, req); err != nil {
			server.RespondWithError(w, err)
			return
		}

		emp, err := h.svc.UpdateEmployee(r.Context(), empID, req)
		if err != nil {
			server.RespondWithError(w, err)
			return
		}
		server.Respond(w, http.StatusOK, emp)
	}
}

func (h *handler) delete() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" {
			server.RespondWithError(w, fault.New(fault.ErrBadRequest, "id not provided", nil))
			return
		}
		empID, err := strconv.Atoi(id)
		if err != nil {
			server.RespondWithError(w, fault.New(fault.ErrBadRequest, "invalid id provided", nil))
			return
		}
		err = h.svc.DeleteEmployee(r.Context(), empID)
		if err != nil {
			server.RespondWithError(w, err)
			return
		}
		server.Respond(w, http.StatusNoContent, nil)
	}
}
