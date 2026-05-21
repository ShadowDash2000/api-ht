package department

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	internalerrors "rest-api-hitalent/internal/errors"
	"rest-api-hitalent/internal/handler/department/dto"
	"rest-api-hitalent/internal/model"
	"strconv"
)

type Service interface {
	Create(ctx context.Context, name string, parentId uint) (model.Department, error)
	Get(ctx context.Context, id uint, depth uint, includeEmployees bool) (*model.DepartmentDetail, error)
	Delete(ctx context.Context, id uint, mode model.DepartmentDeleteMode, reassignToId uint) error
	Update(ctx context.Context, id uint, name *string, parentId *uint) (model.Department, error)
}

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{service: service}
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	var request dto.CreateDepartmentRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	department, err := h.service.Create(r.Context(), request.Name, request.ParentId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(dto.ToCreateDepartmentResponse(department))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	departmentIdStr := r.PathValue("id")
	departmentId, err := strconv.ParseUint(departmentIdStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid department ID", http.StatusBadRequest)
		return
	}

	var depth uint64
	if r.URL.Query().Get("depth") != "" {
		depth, err = strconv.ParseUint(r.URL.Query().Get("depth"), 10, 64)
		if err != nil {
			http.Error(w, "Invalid depth", http.StatusBadRequest)
			return
		}
	} else {
		depth = 1
	}

	var includeEmployees bool
	if r.URL.Query().Get("includeEmployees") != "" {
		includeEmployees, err = strconv.ParseBool(r.URL.Query().Get("include_employees"))
		if err != nil {
			http.Error(w, "Invalid include_employees", http.StatusBadRequest)
			return
		}
	} else {
		includeEmployees = true
	}

	detail, err := h.service.Get(r.Context(), uint(departmentId), uint(depth), includeEmployees)
	if err != nil {
		if errors.Is(err, internalerrors.ErrDepartmentNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(dto.ToGetDepartmentResponse(*detail))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid department ID", http.StatusBadRequest)
		return
	}

	mode, err := model.NewDepartmentDeleteMode(r.URL.Query().Get("mode"))
	if err != nil {
		http.Error(w, "Invalid mode", http.StatusBadRequest)
		return
	}

	var reassignToId uint64
	if mode == model.DepartmentDeleteModeReassign {
		reassignToIdStr := r.URL.Query().Get("reassign_to_department_id")
		if reassignToIdStr == "" {
			http.Error(w, "reassign_to_department_id is required for reassign mode", http.StatusBadRequest)
			return
		}

		reassignToId, err = strconv.ParseUint(reassignToIdStr, 10, 64)
		if err != nil {
			http.Error(w, "Invalid reassign_to_department_id", http.StatusBadRequest)
			return
		}
	}

	err = h.service.Delete(r.Context(), uint(id), mode, uint(reassignToId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *handler) Patch(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid department ID", http.StatusBadRequest)
		return
	}

	var request dto.UpdateDepartmentRequest
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	department, err := h.service.Update(r.Context(), uint(id), request.Name, request.ParentId)
	if err != nil {
		if errors.Is(err, internalerrors.ErrDepartmentSelfParent) || errors.Is(err, internalerrors.ErrDepartmentCycle) {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
		if errors.Is(err, internalerrors.ErrDepartmentNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(dto.ToCreateDepartmentResponse(department))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
