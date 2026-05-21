package employee

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	internalerrors "rest-api-hitalent/internal/errors"
	"rest-api-hitalent/internal/handler/employee/dto"
	"rest-api-hitalent/internal/model"
	"strconv"
)

type Service interface {
	Create(ctx context.Context, data model.NewEmployeeData) (model.Employee, error)
}

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{service: service}
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	departmentIdStr := r.PathValue("id")
	departmentId, err := strconv.ParseUint(departmentIdStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid department ID", http.StatusBadRequest)
		return
	}

	var request dto.CreateEmployeeRequest

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	employee, err := h.service.Create(r.Context(), model.NewEmployeeData{
		DepartmentId: uint(departmentId),
		FullName:     request.FullName,
		Position:     request.Position,
		HiredAt:      request.HiredAt,
	})
	if err != nil {
		if errors.Is(err, internalerrors.ErrDepartmentNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(dto.ToCreateEmployeeResponse(employee))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
